package config

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error){
	encodedPassword := url.QueryEscape(Config.DbPassword)
	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", Config.DbUsername, encodedPassword, Config.DbHost, Config.DbPort, Config.DbName)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err!= nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(Config.DbMaxOpenConnection)
	sqlDB.SetMaxIdleConns(Config.DbMaxIdleConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(Config.DbMaxLifetimeConnection) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(Config.DbMaxIddleTime) * time.Second)

	return db, nil
}