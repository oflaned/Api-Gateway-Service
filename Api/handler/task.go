package handler

import (
	"Mehmat/structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Send(c *gin.Context) {
	code := c.PostForm("code")
	lang := c.PostForm("lang")
	stdIn := c.PostForm("stdIn")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Code is empty"})
		return
	}
	if lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "lang is empty"})
		return
	}

	program := structs.Program{Code: code, Language: lang, StdIn: stdIn}
	out, err := h.services.CompileProgram.RunProgram(program)
	if err != nil {
		fmt.Println(err)
		//Some answer to client here if something bad
		return
	}

	c.JSON(200, gin.H{"message": &out})
}
