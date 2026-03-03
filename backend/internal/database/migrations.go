package database

import (
	"fmt"
	"log"
	"pm-with-essence/config"
	"pm-with-essence/internal/domain/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb(cfg config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	AutoMigration(db)

	return db, nil
}
func IsDatabaseExist(db *gorm.DB, dbName string) error {
	err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)).Error
	if err != nil {
		log.Fatal("Failed to create database. \n", err)
		return err
	}
	return nil
}
func AutoMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{})
	if err != nil {
		log.Fatal("Failed to migrate database. \n", err)
	}
}
