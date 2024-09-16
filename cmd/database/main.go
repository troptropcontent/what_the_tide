package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/gommon/log"
	"github.com/troptropcontent/what_the_tide/config"
)

func main() {
	createMigrationCmd := flag.NewFlagSet("createMigration", flag.ExitOnError)
	createMigrationCmdName := createMigrationCmd.String("name", "", "name")

	forceCmd := flag.NewFlagSet("force", flag.ExitOnError)
	forceVersion := forceCmd.Int("version", 0, "schema version id to wich we want to force the migration back")
	// migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'migrate' or 'create_migration' subcommands")
		os.Exit(1)
	}

	db_migrations_path := config.Root() + "/database/migrations"
	db_file_path := config.Root() + "/database/database.db"

	switch os.Args[1] {

	case "migrate":
		log.Info("Running database migrations")
		m, err := migrate.New(
			"file://"+db_migrations_path,
			"sqlite://"+db_file_path,
		)
		if err != nil {
			log.Fatal("error while creating the migrator instance: ", err)
		}
		err = m.Up()
		if err != nil {
			log.Fatal("error while running the migrations: ", err)
		}
		versionId, _, _ := m.Version()
		log.Infof("migrations successfuly executed, database schema is now at version %d ✅", versionId)
	case "rollback":
		log.Info("Rolling back migrations")
		m, err := migrate.New(
			"file://"+db_migrations_path,
			"sqlite://"+db_file_path,
		)
		if err != nil {
			log.Fatal("error while creating the migrator instance: ", err)
		}
		err = m.Steps(-1)
		if err != nil {
			log.Fatal("error while running the migrations: ", err)
		}
		versionId, _, _ := m.Version()
		log.Infof("migrations successfuly rolleback, database schema is now at version %d ✅", versionId)
	case "create_migration":
		log.Info("Creating new migration files")
		createMigrationCmd.Parse(os.Args[2:])
		if *createMigrationCmdName == "" {
			log.Fatalf("the name flag is required for create_migration subcommand")
		}
		timeStamp := time.Now().UTC().Unix()
		for _, direction := range []string{"up", "down"} {
			migration_file := fmt.Sprintf("%s/%d_%s.%s.sql", db_migrations_path, timeStamp, *createMigrationCmdName, direction)
			err := os.WriteFile(migration_file, []byte(""), 0755)
			if err != nil {
				log.Fatalf("unable to write file: %w", err)
			}
			log.Infof("%s generated ✅")
		}

	case "force":
		m, err := migrate.New(
			"file://"+db_migrations_path,
			"sqlite://"+db_file_path,
		)
		if err != nil {
			log.Fatalf("error creating migrator instance: ", err)
		}
		forceCmd.Parse(os.Args[2:])
		if *forceVersion == 0 {
			log.Fatalf("the version flag is required for force subcommand")
		}

		err = m.Force(*forceVersion)
		if err != nil {
			log.Fatalf("error while running the force operation: ", err)
		}
		log.Infof("database schema forced to version %d ✅", *forceVersion)
	default:
		log.Fatal("expected 'migrate', 'create_migration' or 'force' subcommands")
	}
}
