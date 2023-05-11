package handler

import (
	"Mehmat/lib"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

const bannedFunctions = "new delete strcat strcpy realloc calloc free malloc fopen sprintf fread fwrite fclose fork exec system"

func (h *Handler) Compile(c *gin.Context) {

	code := c.PostForm("code")
	lang := c.PostForm("lang")
	input := c.PostForm("input")

	input = strings.ReplaceAll(input, " ", "\n")
	if !strings.HasSuffix(input, "\n") {
		input += "\n"
	}

	if code == "" {
		log.Print("Code is empty")
		c.String(http.StatusOK, "error: Code is empty")
		return
	}
	if lang == "" {
		log.Print("Lang is empty")
		c.String(http.StatusOK, "error: unknown lang \n")
		return
	}

	if lib.BannedFunctions(code, bannedFunctions) {
		log.Print("Used of banned functions is not allowed")
		c.String(http.StatusOK, "error: Used banned functions \n")
		return
	}

	data := map[string]string{
		"code":  code,
		"lang":  lang,
		"input": input,
	}
	program, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
		c.String(http.StatusOK, "error: Server Error\n")
		return
	}

	out := h.services.RunProgram(program)
	c.String(http.StatusOK, out)
}

func (h *Handler) Program(c *gin.Context) {
	c.HTML(http.StatusOK, "online-compiler.tmpl", gin.H{})
}
