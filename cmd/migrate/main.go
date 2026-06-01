package main

import (
	"log"

	"github.com/MatheusGardin/VReps/internal/infrastructure/db"
	"github.com/MatheusGardin/VReps/internal/infrastructure/db/migrations"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.Init()

	log.Println("Running migrations...")
	if err := migrations.MigrateModels(); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully.")
}
