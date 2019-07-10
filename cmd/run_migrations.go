package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	
	_ "github.com/lib/pq" // Driver
	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
)

const (
	host         = "localhost"
	port         = "5432"
	user         = "jcho"
	password     = "jcho"
	database     = "sling"
	ssl          = "sslmode=disable"
	migrationDir = "file://./migrations/"
)

var log = logrus.WithFields(logrus.Fields{
	"pkg": "cmd",
})

var pgAddr = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", user, password, host, port, database, ssl)

// RunMigrationsCmd is a command to run migration.
var RunMigrationsCmd = &cobra.Command{
	Use:   "runmigrations",
	Short: "run migration on database",
	RunE:  runMigrations,
}

func runMigrations(cmd *cobra.Command, args []string) error {
	migration, err := migrate.New(migrationDir, pgAddr)
	if err != nil {
		return err
	}

	log.Info("performing reset on database")
	if err = migration.Drop(); err != nil {
		return err
	}

	if err := migration.Up(); err != nil {
		return err
	}

	log.Info("migration has been performed successfully")
	return nil
}