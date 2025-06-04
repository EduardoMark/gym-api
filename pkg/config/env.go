package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBTimezone string
	JWTSecret  string
}

func LoadEnv() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error on loading .env")
	}

	return &Env{
		Port:       getEnv("PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "gymapi"),
		DBTimezone: getEnv("DB_TIMEZONE", "America/Sao_Paulo"),
		JWTSecret:  getEnv("JWT_SECRET", "secret_key"),
	}
}

func getEnv(key, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	return fallback
}
