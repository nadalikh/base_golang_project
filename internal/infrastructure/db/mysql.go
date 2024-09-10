package db

import (
	"fmt"
	"log"
	"testGorm/internal/domain/entity"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Set timezone
	time.LoadLocation("Asia/Tehran")

	// Migrate the schema
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate: %w", err)
	}

	log.Println("Database connection and migration successful")
	return db, nil
}
