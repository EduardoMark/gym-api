package main

import (
	"github.com/EduardoMark/gym-api/internal/api/handler"
	"github.com/EduardoMark/gym-api/internal/equipament"
	"github.com/EduardoMark/gym-api/internal/user"
	"github.com/EduardoMark/gym-api/pkg/auth"
	"github.com/EduardoMark/gym-api/pkg/config"
	"github.com/EduardoMark/gym-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadEnv()
	database.ConnectPostgres(*cfg)
	database.AutoMigrate(&user.User{}, &equipament.Equipament{})

	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	userRepo := user.NewRepository(database.DB)
	userUseCase := user.NewUserUseCase(userRepo, auth.NewJWTAuth(cfg.JWTSecret))
	userHandler := handler.NewUserHandler(userUseCase)
	userHandler.RegisterRoutes(apiV1)

	authHandler := handler.NewAuthHandler(userUseCase)
	authHandler.RegisterRouter(apiV1)

	equipamentRepo := equipament.NewRepository(database.DB)
	equipamentUsecase := equipament.NewEquipamentUseCase(equipamentRepo)
	equipamentHandler := handler.NewEquipamentHandler(equipamentUsecase)
	equipamentHandler.RegisterRoutes(apiV1)

	router.Run(":8080")
}
