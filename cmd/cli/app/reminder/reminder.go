//
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

// Package reminder contains the logic for background entity reconciliation
package reminder

import (
	"github.com/spf13/cobra"
	"github.com/stacklok/minder/cmd/cli/app"
	ghclient "github.com/stacklok/minder/internal/providers/github"
)

var ReminderCmd = &cobra.Command{
	Use:   "reminder",
	Short: "Manage sending reminders to minder about various entities",
	Long:  `The reminder commands allow to send reminders to minder about various entities. Reminder is equivalent to processing the entities in the background.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Usage()
	},
}

func init() {
	app.RootCmd.AddCommand(ReminderCmd)
	// Add the reminder command to the root command
	// Flags for all subcommands
	ReminderCmd.PersistentFlags().StringP("provider", "p", ghclient.Github, "Name of the provider, e.g. github")
	ReminderCmd.PersistentFlags().StringP("project", "j", "", "ID of the project")
}
