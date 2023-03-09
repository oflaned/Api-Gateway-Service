package handler

import (
	"Mehmat/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Compile(c *gin.Context) {
	code := c.PostForm("code")
	lang := c.PostForm("lang")
	stdIn := c.PostForm("stdIn")

	if code == "" {
		c.HTML(http.StatusBadRequest, "err.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": "Code is empty",
		})
		return
	}
	if lang == "" {
		c.HTML(http.StatusBadRequest, "err.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": "Lang is empty",
		})
		return
	}

	program := structs.Program{Code: code, Language: lang, StdIn: stdIn}
	out, err := h.services.CompileProgram.RunProgram(program)
	if err != nil {
		c.HTML(http.StatusBadRequest, "err.tmpl", gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": &out})
}
