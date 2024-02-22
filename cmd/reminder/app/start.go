package app

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stacklok/minder/internal/config"
	reminderconfig "github.com/stacklok/minder/internal/config/reminder"
	"github.com/stacklok/minder/internal/db"
	"github.com/stacklok/minder/internal/reminder"
	"golang.org/x/sync/errgroup"
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

	// Database configuration
	dbConn, _, err := cfg.Database.GetDBConnection(ctx)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Printf("error closing database connection: %v", err)
		}
	}(dbConn)

	store := db.NewStore(dbConn)
	reminder := reminder.NewReminder(store, cfg)
	defer reminder.Stop()

	errg, ctx := errgroup.WithContext(ctx)

	errg.Go(func() error {
		return reminder.Start(ctx)
	})

	return errg.Wait()
}
