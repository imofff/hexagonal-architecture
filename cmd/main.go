package main

import (
	"hexagonal/internal/adapter/http"
	"hexagonal/internal/adapter/http/handler"
	"hexagonal/internal/adapter/postgres"
	"hexagonal/internal/app/usecase"
	"hexagonal/internal/config"
)

func main() {
	db := config.NewPostgresConnection()

	repo := postgres.NewUserRepo(db)
	usecase := usecase.NewUserUsecase(repo)
	userHandler := handler.NewUserHandler(usecase)

	r := http.NewRouter(userHandler)

	// Run server
	config.StartServer(r) // start server (optional wrap)
}
