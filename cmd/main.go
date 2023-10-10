package main

import (
	"greedy/pkg/delivery"
	"greedy/pkg/domain"
	"greedy/pkg/repository"
	"greedy/pkg/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	Config, err := domain.LoadConfig()
	if err != nil {
		log.Fatalln("Error loading config : ", err)
	}

	server, err := InitializeApi(Config)
	if err != nil {
		log.Fatalln("Error initializing server : ", err)
	}

	server.Run(Config.Port)
}

func InitializeApi(config *domain.Config) (*gin.Engine, error) {
	router := gin.Default()

	userRepo := repository.NewRepoLayer(config)

	userUsecase := usecase.NewUsecaseLayer(config, userRepo)

	userHandlers := delivery.NewHandlers(config, userUsecase)

	userRoutes := delivery.NewRoutes(config, userHandlers)

	userRoutes.SetKeyValueRoutes(router)
	userRoutes.SetQueRoutes(router)

	return router, nil
}
