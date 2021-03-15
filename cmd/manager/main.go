package main

import (
	"github.com/joho/godotenv"
	"github/kaanaktas/task-manager/internal/config"
	"github/kaanaktas/task-manager/internal/store"
	"github/kaanaktas/task-manager/pkg/manager"
	"log"
	"os"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := config.NewEchoEngine()
	dbx := store.LoadDBConnection()

	managerRepository := manager.NewRepository(dbx)
	managerService := manager.NewService(managerRepository)

	manager.RegisterHandler(e, managerService)

	log.Printf("starting server at :%s", port)

	if err := e.Start(":" + port); err != nil {
		log.Fatalf("error while starting server at :%s, %v", port, err)
	}
}
