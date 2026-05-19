package main

import (
	"log"
	"os"

	db "github.com/ElCesarSP/go-api/prisma/db"

	"github.com/ElCesarSP/go-api/internal/controller"
	"github.com/ElCesarSP/go-api/internal/repository"
	"github.com/ElCesarSP/go-api/internal/routes"
	"github.com/ElCesarSP/go-api/internal/service"
)

func main() {
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer client.Prisma.Disconnect()

	userRepo := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := routes.Initialize(userController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
