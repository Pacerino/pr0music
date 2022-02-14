package database

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func init() {
	for _, env := range []string{"DB_HOST", "DB_USER", "DB_PASS", "DB_DATABASE", "DB_PORT", "DB_SSL"} {
		if len(os.Getenv(env)) == 0 {
			log.Fatal(fmt.Sprintf("Missing %s from environment", env))
		}
	}
}

func Connect() (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sql, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = sql.Ping()
	if err != nil {
		return nil, err
	}
	log.Info("Connected to Database")

	if err := db.AutoMigrate(&Items{}, &Comments{}); err != nil {
		log.WithError(err).Fatal("Could not migrate Items and/or Comments")
	} else {
		log.Debug("Auto Migrated Items and Comments")
	}

	return &DB{db}, nil
}
