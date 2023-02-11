package main

import (
	"Mehmat/lib"
	"log"
	"os"
)

var projectPath string

func init() {
	var err error
	projectPath, err = os.Getwd()
	if err != nil {
		log.Fatal("Project path can't be declared")
	}
	log.Println("Project Path: ", projectPath)
}

func main() {

	//Тестовый билд программы
	var binFolder = "/tasks"
	_, err := lib.BuildGoApp(projectPath, binFolder)
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

}
