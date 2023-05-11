package main

import (
	"Mehmat/Api/handler"
	"Mehmat/Api/service"
	"Mehmat/config"
	"Mehmat/utils"
	"log"
)

func main() {
	cfg := config.GetConfig()

	log.Println(cfg.Port)

	dep, err := utils.CheckDependencies()
	if err != nil {
		log.Fatalf("dependency %s not found: %v", dep, err)
	}

	services := service.NewService()
	handlers := handler.NewHandler(services)

	server := new(Server)
	err = server.Run(cfg.Port, handlers.InitRoutes())
	if err != nil {
		log.Fatal("Server are not run")
	}

}
