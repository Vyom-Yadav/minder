// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package pull_request provides the pull request remediation engine
package pull_request

import (
	"bytes"
	"context"
	"crypto/sha1" // #nosec G505 - we're not using sha1 for crypto, only to quickly compare contents
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5/plumbing/filemode"
	"github.com/rs/zerolog/log"
	fzconfig "github.com/stacklok/frizbee/pkg/config"
	"github.com/stacklok/frizbee/pkg/ghactions"
	"github.com/stacklok/frizbee/pkg/utils"
	"gopkg.in/yaml.v3"

	"github.com/stacklok/minder/internal/engine/interfaces"
	"github.com/stacklok/minder/internal/util"
	pb "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
	v1 "github.com/stacklok/minder/pkg/providers/v1"
)

// See minder.proto for more detailed documentation
const (
	// minderContentModification replaces the contents of a file with the given template
	minderContentModification = "minder.content"
	// minderFrizbeeTagResolve replaces a github action tag with the appropriate checksum
	minderFrizbeeTagResolve = "minder.actions.replace_tags_with_sha"
)

type fsEntry struct {
	contentTemplate *template.Template

	Path    string `json:"path"`
	Content string `json:"content"`
	Mode    string `json:"mode"`
}

func (fe *fsEntry) write(fs billy.Filesystem) error {
	dirOsMode, err := filemode.Dir.ToOSFileMode()
	if err != nil {
		return fmt.Errorf("cannot get directory mode: %w", err)
	}
	if err := fs.MkdirAll(filepath.Dir(fe.Path), dirOsMode); err != nil {
		return fmt.Errorf("cannot create directory: %w", err)
	}

	parsedGitMode, err := filemode.New(fe.Mode)
	if err != nil {
		return fmt.Errorf("cannot parse mode: %w", err)
	}
	parsedOsMode, err := parsedGitMode.ToOSFileMode()
	if err != nil {
		return fmt.Errorf("cannot convert file mode: %w", err)
	}

	f, err := fs.OpenFile(fe.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, parsedOsMode)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close file")
		}
	}()

	_, err = io.WriteString(f, fe.Content)
	if err != nil {
		return fmt.Errorf("cannot write to file: %w", err)
	}

	return nil
}

type fsChangeSet struct {
	fs      billy.Filesystem
	entries []*fsEntry
}

func (fcs *fsChangeSet) writeEntries() error {
	for i := range fcs.entries {
		entry := fcs.entries[i]

		if err := entry.write(fcs.fs); err != nil {
			return fmt.Errorf("cannot write entry %s: %w", entry.Path, err)
		}
	}

	return nil
}

func (fcs *fsChangeSet) hash() (string, error) {
	if fcs.entries == nil {
		return "", fmt.Errorf("no entries")
	}

	var combinedContents string

	for i := range fcs.entries {
		if len(fcs.entries[i].Content) == 0 {
			// just making sure we call sha1() after expandContents()
			return "", fmt.Errorf("content (index %d) is empty", i)
		}
		combinedContents += fcs.entries[i].Path + fcs.entries[i].Content
	}

	// #nosec G401 - we're not using sha1 for crypto, only to quickly compare contents
	return fmt.Sprintf("%x", sha1.Sum([]byte(combinedContents))), nil
}

func (fcs *fsChangeSet) writeSummary(out io.Writer) error {
	if fcs.entries == nil {
		return fmt.Errorf("no entries")
	}

	b, err := json.MarshalIndent(fcs.entries, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal entries: %w", err)
	}
	fmt.Fprintln(out, b)

	return nil
}

type fsModifier interface {
	hash() (string, error)
	writeSummary(out io.Writer) error
	createFsModEntries(ctx context.Context, params interfaces.ActionsParams) error
	modifyFs() ([]*fsEntry, error)
}

type modificationConstructorParams struct {
	prCfg *pb.RuleType_Definition_Remediate_PullRequestRemediation
	ghCli v1.GitHub
	bfs   billy.Filesystem
}

type modificationConstructor func(*modificationConstructorParams) (fsModifier, error)

type modificationRegistry map[string]modificationConstructor

func newModificationRegistry() modificationRegistry {
	return make(map[string]modificationConstructor)
}

func (mr modificationRegistry) register(name string, constructor modificationConstructor) {
	mr[name] = constructor
}

func (mr modificationRegistry) registerBuiltIn() {
	mr.register(minderContentModification, newContentModification)
	mr.register(minderFrizbeeTagResolve, newFrizbeeTagResolveModification)
}

func (mr modificationRegistry) getModification(
	name string,
	params *modificationConstructorParams,
) (fsModifier, error) {
	constructor, ok := mr[name]
	if !ok {
		return nil, fmt.Errorf("unknown modification type: %s", name)
	}

	return constructor(params)
}

var _ fsModifier = (*contentModification)(nil)

type contentModification struct {
	fsChangeSet
	prCfg *pb.RuleType_Definition_Remediate_PullRequestRemediation
}

var _ modificationConstructor = newContentModification

func newContentModification(
	params *modificationConstructorParams,
) (fsModifier, error) {
	// validate
	if params.prCfg == nil {
		return nil, fmt.Errorf("pull request config cannot be nil")
	}

	if len(params.prCfg.Contents) == 0 {
		return nil, fmt.Errorf("pull request config contents cannot be empty")
	}

	for _, cnt := range params.prCfg.Contents {
		if cnt.Path == "" {
			return nil, fmt.Errorf("pull request config contents path cannot be empty")
		}
		if cnt.Content == "" {
			return nil, fmt.Errorf("pull request config contents content cannot be empty")
		}
	}

	entries, err := prConfigToEntries(params.prCfg)
	if err != nil {
		return nil, fmt.Errorf("cannot create PR entries: %w", err)
	}

	return &contentModification{
		prCfg: params.prCfg,
		fsChangeSet: fsChangeSet{
			entries: entries,
			fs:      params.bfs,
		},
	}, nil
}

func prConfigToEntries(prCfg *pb.RuleType_Definition_Remediate_PullRequestRemediation) ([]*fsEntry, error) {
	entries := make([]*fsEntry, len(prCfg.Contents))
	for i := range prCfg.Contents {
		cnt := prCfg.Contents[i]

		contentTemplate, err := util.ParseNewTextTemplate(&cnt.Content, fmt.Sprintf("Content[%d]", i))
		if err != nil {
			return nil, fmt.Errorf("cannot parse content template (index %d): %w", i, err)
		}

		mode := ghModeNonExecFile
		if cnt.GetMode() != "" {
			mode = *cnt.Mode
		}

		entries[i] = &fsEntry{
			Path:            cnt.Path,
			Mode:            mode,
			contentTemplate: contentTemplate,
		}
	}

	return entries, nil
}

func (ca *contentModification) createFsModEntries(
	_ context.Context,
	params interfaces.ActionsParams,
) error {
	data := map[string]interface{}{
		"Params":  params.GetRule().Params.AsMap(),
		"Profile": params.GetRule().Def.AsMap(),
	}
	for i, entry := range ca.entries {
		content := new(bytes.Buffer)

		if err := entry.contentTemplate.Execute(content, data); err != nil {
			return fmt.Errorf("cannot execute content template (index %d): %w", i, err)
		}
		entry.Content = content.String()
	}

	return nil

}

func (ca *contentModification) modifyFs() ([]*fsEntry, error) {
	err := ca.fsChangeSet.writeEntries()
	if err != nil {
		return nil, fmt.Errorf("cannot write entries: %w", err)
	}

	return ca.entries, nil
}

var _ fsModifier = (*frizbeeTagResolveModification)(nil)

type frizbeeTagResolveModification struct {
	fsChangeSet

	ghCli v1.GitHub
}

var _ modificationConstructor = newFrizbeeTagResolveModification

func newFrizbeeTagResolveModification(
	params *modificationConstructorParams,
) (fsModifier, error) { // nolint:unparam // we need to match the interface
	return &frizbeeTagResolveModification{
		fsChangeSet: fsChangeSet{
			fs: params.bfs,
		},
		ghCli: params.ghCli,
	}, nil
}

func (ftr *frizbeeTagResolveModification) createFsModEntries(ctx context.Context, _ interfaces.ActionsParams) error {
	entries := []*fsEntry{}

	err := ghactions.TraverseGitHubActionWorkflows(ftr.fs, ".github/workflows", func(path string, wflow *yaml.Node) error {
		m, err := ghactions.ModifyReferencesInYAML(ctx, ftr.ghCli, wflow, &fzconfig.GHActions{})
		if err != nil {
			return fmt.Errorf("failed to process YAML file %s: %w", path, err)
		}

		buf, err := utils.YAMLToBuffer(wflow)
		if err != nil {
			return fmt.Errorf("failed to convert YAML to buffer: %w", err)
		}

		if m {
			entries = append(entries, &fsEntry{
				Path:    path,
				Content: buf.String(),
				Mode:    filemode.Regular.String(),
			})
		}

		return nil
	})
	if err != nil {
		return err
	}

	ftr.entries = entries
	return nil
}

func (ftr *frizbeeTagResolveModification) modifyFs() ([]*fsEntry, error) {
	err := ftr.fsChangeSet.writeEntries()
	if err != nil {
		return nil, fmt.Errorf("cannot write entries: %w", err)
	}
	return ftr.entries, nil
}