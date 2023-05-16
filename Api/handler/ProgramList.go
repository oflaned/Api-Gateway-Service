package handler

import (
	"Mehmat/model/program"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Squares struct {
	Language string
	Code     string
	Name     string
}

func (h *Handler) AddProgram(c *gin.Context) {

	code := c.PostForm("code")
	lang := c.PostForm("lang")
	name := c.PostForm("name")

	prog := program.Program{Name: name, Lang: lang, Code: code}

	id, err := h.services.Repository.Create(context.Background(), prog)
	if err != nil {
		//Error to page
	}
	log.Printf("Add new Program to db. ID:%s", id)

}

func (h *Handler) ProgramsList(c *gin.Context) {
	programs, err := h.services.Repository.FindAll(context.Background())
	if err != nil {
		c.HTML(http.StatusInternalServerError, "err.tmpl", gin.H{})
	}

	squares := make([]Squares, 0)
	for _, p := range programs {
		square := Squares{
			Language: p.Lang,
			Code:     p.Code,
			Name:     p.Name,
		}
		squares = append(squares, square)
	}

	c.HTML(http.StatusOK, "programs.tmpl", gin.H{
		"blocks": squares,
	})
}
