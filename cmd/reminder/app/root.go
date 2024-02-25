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

// Package app provides the cli subcommands for managing a minder control plane
package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/stacklok/minder/internal/config"
	reminderconfig "github.com/stacklok/minder/internal/config/reminder"
	"github.com/stacklok/minder/internal/util/cli"
)

var (
	// RootCmd represents the base command when called without any subcommands
	RootCmd = &cobra.Command{
		Use:   "re-minder",
		Short: "Re-Minder process for sending reminders to minder server",
		Long:  `re-minder sends reminders to the minder server to process entities in background.`,
	}
)

const configFileName = "reminder-config.yaml"

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	RootCmd.SetOut(os.Stdout)
	RootCmd.SetErr(os.Stderr)
	err := RootCmd.ExecuteContext(context.Background())
	cli.ExitNicelyOnError(err, "Error executing root command")
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().String("config", "", fmt.Sprintf("config file (default is $PWD/%s)", configFileName))
	if err := viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config")); err != nil {
		RootCmd.Printf("error: %s", err)
		os.Exit(1)
	}
}

func initConfig() {
	cfgFile := viper.GetString("config")
	cfgFileData, err := config.GetConfigFileData(cfgFile, filepath.Join(".", configFileName))
	if err != nil {
		RootCmd.PrintErrln(err)
		os.Exit(1)
	}

	keysWithNullValue := config.GetKeysWithNullValueFromYAML(cfgFileData, "")
	if len(keysWithNullValue) > 0 {
		RootCmd.PrintErrln("Error: The following configuration keys are missing values:")
		for _, key := range keysWithNullValue {
			RootCmd.PrintErrln("Null Value at: " + key)
		}
		os.Exit(1)
	}

	// TODO: Test whether this works or not, is this ok?
	cfg, ok := cfgFileData.(*reminderconfig.Config)
	if !ok {
		RootCmd.Printf("invalid config type: %T\n", cfgFileData)
		os.Exit(1)
	}

	err = reminderconfig.ValidateConfig(cfg)
	if err != nil {
		RootCmd.PrintErrln(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// use defaults
		viper.SetConfigName(strings.TrimSuffix(configFileName, filepath.Ext(configFileName)))
		viper.AddConfigPath(".")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
	}
}
