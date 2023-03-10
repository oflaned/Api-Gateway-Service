package handler

import (
	"Mehmat/structs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) Compile(c *gin.Context) {
	code := c.PostForm("code")
	lang := c.PostForm("lang")
	input := c.PostForm("input")

	if code == "" {
		log.Print("Code is empty")
		c.HTML(http.StatusBadRequest, "err.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": "Code is empty",
		})
		return
	}
	if lang == "" {
		log.Print("Lang is empty")
		c.HTML(http.StatusBadRequest, "err.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": "Lang is empty",
		})
		return
	}

	program := structs.Program{Code: code, Language: lang, StdIn: input}
	out, err := h.services.CompileProgram.RunProgram(program)
	if err != nil {
		log.Print(err)
		c.HTML(http.StatusInternalServerError, "err.tmpl", gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Server Error",
		})
		return
	}

	c.String(http.StatusOK, out)
}

func (h *Handler) Program(c *gin.Context) {
	c.HTML(http.StatusOK, "program.tmpl", gin.H{})
}
