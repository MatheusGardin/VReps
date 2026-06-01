package main

import (
	"log"

	"github.com/MatheusGardin/VReps/cmd/api/di"
	apiRunner "github.com/MatheusGardin/VReps/internal/infrastructure/api"
	"github.com/MatheusGardin/VReps/internal/infrastructure/common"
	"github.com/MatheusGardin/VReps/internal/infrastructure/db"
	"github.com/MatheusGardin/VReps/internal/infrastructure/db/migrations"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.Init()
	if err := migrations.MigrateModels(); err != nil {
		log.Fatalf("failed to run database migrations: %v", err)
	}

	api := di.InitializeServer()

	router := api.Router.RegisterRoutes()

	apiRunner.Run(router)

	common.WaitOsInterruption()
}
