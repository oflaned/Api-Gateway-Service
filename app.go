package main

import (
	"Mehmat/Api/handler"
	"Mehmat/Api/repository"
	service "Mehmat/Api/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(Server)
	err := server.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatal("Server are not running")
	}

}

/*
	//Тестовый билд программы
	var binFolder = "/tasks"
	_, err = lib.BuildGoApp(projectPath, binFolder)
	if err != nil {
		log.Println(err)
	}

	//Тестовый запуск программы
	var input = "Hey\nMax\n1\n2\n"
	out, err := lib.RunApp(projectPath+binFolder+"/tasks", input)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(out)
	}
*/