package main

import (
	"Mehmat/Api/handler"
	"Mehmat/Api/service"
	"Mehmat/config"
	"Mehmat/model/mossStatus"
	"Mehmat/model/mossStatus/db"
	"Mehmat/model/program"
	programDB "Mehmat/model/program/db"
	"Mehmat/pkg/client/postgressql"
	"Mehmat/utils"
	"context"
	"log"
)

func main() {
	var repPrograms program.Repository
	var repStatus mossStatus.Repository
	var postgresSQLClient postgressql.Client
	var err error

	cfg := config.GetConfig()
	if dep, err := utils.CheckDependencies(); err != nil {
		log.Fatalf("dependency %s not found: %v", dep, err)
	}

	if postgresSQLClient, err = postgressql.NewClient(context.TODO(), cfg.Storage); err != nil {
		log.Fatalf("Error while create db client")
	}

	repPrograms = programDB.NewRepository(postgresSQLClient)
	repStatus = db.NewRepository(postgresSQLClient)
	services := service.NewService(repPrograms, repStatus)
	handlers := handler.NewHandler(services)

	server := new(Server)
	if server.Run(cfg.Port, handlers.InitRoutes()); err != nil {
		log.Fatal("Server are not run")
	}

}
