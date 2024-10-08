// Copyright 2024 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package manager contains the GitLabProviderClassManager
package manager

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"slices"

	"github.com/rs/zerolog"

	"github.com/stacklok/minder/internal/config/server"
	"github.com/stacklok/minder/internal/crypto"
	"github.com/stacklok/minder/internal/db"
	"github.com/stacklok/minder/internal/events"
	"github.com/stacklok/minder/internal/providers/credentials"
	"github.com/stacklok/minder/internal/providers/gitlab"
	v1 "github.com/stacklok/minder/pkg/providers/v1"
)

type providerClassManager struct {
	store    db.Store
	crypteng crypto.Engine
	// gitlab provider config
	glpcfg        *server.GitLabConfig
	webhookURL    string
	parentContext context.Context
	pub           events.Publisher

	// secrets for the webhook. These are stored in the
	// structure to allow efficient fetching. Rotation
	// requires a process restart.
	currentWebhookSecret   string
	previousWebhookSecrets []string
}

// NewGitLabProviderClassManager creates a new provider class manager for the dockerhub provider
func NewGitLabProviderClassManager(
	ctx context.Context, crypteng crypto.Engine, store db.Store, pub events.Publisher,
	cfg *server.GitLabConfig, wgCfg server.WebhookConfig,
) (*providerClassManager, error) {
	webhookURLBase := wgCfg.ExternalWebhookURL
	if webhookURLBase == "" {
		return nil, errors.New("webhook URL is required")
	}

	if cfg == nil {
		return nil, errors.New("gitlab config is required")
	}

	webhookURL, err := url.JoinPath(webhookURLBase, url.PathEscape(string(db.ProviderClassGitlab)))
	if err != nil {
		return nil, fmt.Errorf("error joining webhook URL: %w", err)
	}

	whSecret, err := cfg.GetWebhookSecret()
	if err != nil {
		return nil, fmt.Errorf("error getting webhook secret: %w", err)
	}

	previousSecrets, err := cfg.GetPreviousWebhookSecrets()
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("previous secrets not loaded")
	}

	return &providerClassManager{
		store:                  store,
		crypteng:               crypteng,
		pub:                    pub,
		glpcfg:                 cfg,
		webhookURL:             webhookURL,
		parentContext:          ctx,
		currentWebhookSecret:   whSecret,
		previousWebhookSecrets: previousSecrets,
	}, nil
}

// GetSupportedClasses implements the ProviderClassManager interface
func (_ *providerClassManager) GetSupportedClasses() []db.ProviderClass {
	return []db.ProviderClass{db.ProviderClassGitlab}
}

// Build implements the ProviderClassManager interface
func (g *providerClassManager) Build(ctx context.Context, config *db.Provider) (v1.Provider, error) {
	class := config.Class
	// This should be validated by the caller, but let's check anyway
	if !slices.Contains(g.GetSupportedClasses(), class) {
		return nil, fmt.Errorf("provider does not implement gitlab")
	}

	if config.Version != v1.V1 {
		return nil, fmt.Errorf("provider version not supported")
	}

	creds, err := g.getProviderCredentials(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch credentials")
	}

	cfg, err := gitlab.ParseV1Config(config.Definition)
	if err != nil {
		return nil, fmt.Errorf("error parsing gitlab config: %w", err)
	}

	cli, err := gitlab.New(creds, cfg, g.webhookURL, g.currentWebhookSecret)
	if err != nil {
		return nil, fmt.Errorf("error creating gitlab client: %w", err)
	}
	return cli, nil
}

// Delete implements the ProviderClassManager interface
// TODO: Implement this
func (_ *providerClassManager) Delete(_ context.Context, _ *db.Provider) error {
	return nil
}

func (m *providerClassManager) getProviderCredentials(
	ctx context.Context,
	prov *db.Provider,
) (v1.GitLabCredential, error) {
	encToken, err := m.store.GetAccessTokenByProjectID(ctx,
		db.GetAccessTokenByProjectIDParams{Provider: prov.Name, ProjectID: prov.ProjectID})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("error getting credential: %w", err)
	}

	// TODO: get rid of this once we migrate all secrets to use the new structure
	var encryptedData crypto.EncryptedData
	if encToken.EncryptedAccessToken.Valid {
		encryptedData, err = crypto.DeserializeEncryptedData(encToken.EncryptedAccessToken.RawMessage)
		if err != nil {
			return nil, err
		}
	} else if encToken.EncryptedToken.Valid {
		encryptedData = crypto.NewBackwardsCompatibleEncryptedData(encToken.EncryptedToken.String)
	} else {
		return nil, fmt.Errorf("no secret found for provider %s", encToken.Provider)
	}
	decryptedToken, err := m.crypteng.DecryptOAuthToken(encryptedData)
	if err != nil {
		return nil, fmt.Errorf("error decrypting access token: %w", err)
	}

	return credentials.NewGitLabTokenCredential(decryptedToken.AccessToken), nil
}

func (m *providerClassManager) MarshallConfig(
	_ context.Context, class db.ProviderClass, config json.RawMessage,
) (json.RawMessage, error) {
	if !slices.Contains(m.GetSupportedClasses(), class) {
		return nil, fmt.Errorf("provider does not implement %s", string(class))
	}

	return gitlab.MarshalV1Config(config)
}
