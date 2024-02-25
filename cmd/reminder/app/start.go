package app

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/stacklok/minder/internal/config"
	reminderconfig "github.com/stacklok/minder/internal/config/reminder"
	"github.com/stacklok/minder/internal/reminder"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the re-minder process",
	Long:  `Start the re-minder process to send reminders to the minder server to process entities in background.`,
	RunE:  start,
}

func start(cmd *cobra.Command, _ []string) error {
	ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt)
	defer cancel()

	cfg, err := config.ReadConfigFromViper[reminderconfig.Config](viper.GetViper())
	if err != nil {
		return fmt.Errorf("unable to read config: %w", err)
	}

	err = reminderconfig.ValidateConfig(cfg)
	if err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// TODO: Maybe add a logger?

	reminderService := reminder.NewReminder(cfg)
	defer reminderService.Stop()

	errg, ctx := errgroup.WithContext(ctx)

	errg.Go(func() error {
		return reminderService.Start(ctx)
	})

	return errg.Wait()
}

func init() {
	RootCmd.AddCommand(startCmd)

	if err := reminderconfig.RegisterReminderFlags(viper.GetViper(), RootCmd.PersistentFlags()); err != nil {
		log.Fatal().Err(err).Msg("Error registering reminder flags")
	}
}
