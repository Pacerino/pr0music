package database

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Session struct {
	*gorm.DB
}

func Connect() (*Session, error) {
	for _, env := range []string{"DB_HOST", "DB_USER", "DB_PASS", "DB_DATABASE", "DB_PORT", "DB_SSL"} {
		if len(os.Getenv(env)) == 0 {
			log.Fatal(fmt.Sprintf("Missing %s from environment!", env))
		}
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Berlin",
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
	var date string
	db.Raw("select current_date;").Scan(&date)
	log.Info(fmt.Sprintf("Connected to Database! (Current Date: %s)", date))

	if err := db.AutoMigrate(&Items{}, &Comments{}); err != nil {
		log.WithError(err).Fatal("Could not migrate Items and/or Comments!")
	} else {
		log.Debug("Auto Migrated Items and Comments!")
	}

	return &Session{db}, nil
}
