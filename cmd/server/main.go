package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"go-user-api/config"
	dbsqlc "go-user-api/db/sqlc"
	"go-user-api/internal/handler"
	"go-user-api/internal/middleware"
	"go-user-api/internal/repository"
	"go-user-api/internal/routes"
	"go-user-api/internal/service"
)

func main() {
	app := fiber.New()

	// DB
	pool := config.ConnectDB()
	queries := dbsqlc.New(pool)

	// Layers
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	// Routes
	routes.SetupRoutes(app, h)

	log.Fatal(app.Listen(":3000"))

	app.Use(middleware.RequestMiddleware())
}
