package main

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func connectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "pr0music.", // schema name
			SingularTable: false,
		}})
	if err != nil {
		return nil, err
	}

	sql, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sql.Ping(); err != nil {
		return nil, err
	}
	log.Info("Connected to Database")

	if err := db.AutoMigrate(&Items{}, &Comments{}); err != nil {
		log.WithError(err).Fatal("Could not migrate Items and/or Comments")
		return nil, err
	}
	log.Debug("Auto Migrated Items and Comments")

	return db, nil
}
