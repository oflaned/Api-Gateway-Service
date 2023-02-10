package main

import (
	"Mehmat/lib"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	log.Println("Start initialization")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found")
	}

	log.Println("Finish initialization")
}

func main() {
	log.Println("Start the Program")

	projectPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Project path can't be declared")
	}
	log.Println("Project Path: ", projectPath)

	//Тестовый билд программы
	_, err = lib.BuildGoApp(projectPath + "/tasks")
	if err != nil {
		log.Fatal(err)
	}

	//Тестовый запуск программы
	var input = "Hey\nMax\n1\n2\n"
	out, err := lib.RunApp(projectPath+"/tasks/tasks", input)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out)
}
