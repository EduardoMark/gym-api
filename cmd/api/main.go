package main

import (
	"github.com/EduardoMark/gym-api/pkg/config"
	"github.com/EduardoMark/gym-api/pkg/database"
)

func main() {
	cfg := config.LoadEnv()
	database.ConnectPostgres(*cfg)
}
