package handler

import (
	"Mehmat/Api/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("./templates/**/*")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	compiler := router.Group("/online-compiler")
	{
		compiler.GET("", h.Program)
		compiler.POST("/send", h.Compile)
	}
	programsList := router.Group("/programs")
	{
		programsList.GET("", h.ProgramsList)
		programsList.POST("/send", h.AddProgram)
	}
	antiPlagiat := router.Group("/antiplagiat")
	{
		antiPlagiat.POST("/send", h.SendToMoss)
		antiPlagiat.GET("", h.AntiplagiatStatus)
	}
	return router

}
