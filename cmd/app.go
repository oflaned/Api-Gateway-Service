package main

import (
	OnlineCompiler "Mehmat"
	"Mehmat/Api/handler"
	"Mehmat/Api/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal("error initialization config")
	}

	services := service.NewService()
	handlers := handler.NewHandler(services)

	server := new(OnlineCompiler.Server)
	err = server.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatal("Server are not run")
	}

}
func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
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
