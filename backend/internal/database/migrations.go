package database

import (
	"fmt"
	"log"
	"pm-with-essence/config"
	"pm-with-essence/internal/domain/entity/pm-entity"

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
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	AutoMigration(db)

	return db, nil
}

func AutoMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&pm_entity.User{},
		&pm_entity.Project{},
		&pm_entity.Board{},
		&pm_entity.Column{},
		&pm_entity.Issue{})
	if err != nil {
		log.Fatal("Failed to migrate database. \n", err)
	}

	if err1 := Seed(db); err1 != nil {
		log.Println(err1)
	}
}
