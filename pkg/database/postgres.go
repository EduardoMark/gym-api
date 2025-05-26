package database

import (
	"fmt"
	"log"

	"github.com/EduardoMark/gym-api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres(cfg config.Env) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBTimezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error on connect with postgres database: %v", err)
	}

	DB = db
}

func AutoMigrate(models ...interface{}) {
	DB.AutoMigrate(models...)
}
