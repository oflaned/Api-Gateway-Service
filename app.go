package main

import (
	"Mehmat/Api/handler"
	"Mehmat/Api/service"
	"Mehmat/config"
	"Mehmat/model/program"
	programDB "Mehmat/model/program/db"
	"Mehmat/pkg/client/postgressql"
	"Mehmat/utils"
	"context"
	"log"
)

func main() {
	var repository program.Repository
	var postgresSQLClient postgressql.Client
	var err error

	cfg := config.GetConfig()
	if dep, err := utils.CheckDependencies(); err != nil {
		log.Fatalf("dependency %s not found: %v", dep, err)
	}

	if postgresSQLClient, err = postgressql.NewClient(context.TODO(), cfg.Storage); err != nil {
		log.Fatalf("Error while create db client")
	}

	repository = programDB.NewRepository(postgresSQLClient)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(Server)
	if server.Run(cfg.Port, handlers.InitRoutes()); err != nil {
		log.Fatal("Server are not run")
	}

}
