package handler

import (
	"Mehmat/model/mossStatus"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
	"io/ioutil"
	"log"
	"net/http"
)

type SquaresStatus struct {
	Date   pgtype.Date
	Status string
	Link   string
}

func (h *Handler) SendToMoss(c *gin.Context) {
	var selectedPrograms []struct {
		ID int `json:"id"`
	}

	if err := c.BindJSON(&selectedPrograms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка чтения данных"})
		return
	}

	programs, err := h.services.ProgramRep.FindAll(context.Background())
	if err != nil {
		log.Print(err)
		c.HTML(http.StatusInternalServerError, "err.tmpl", gin.H{})
		return
	}

	programList := make([]map[string]interface{}, 0)
	for _, program := range selectedPrograms {
		programData := map[string]interface{}{
			"name": programs[program.ID-1].Name,
			"lang": programs[program.ID-1].Lang,
			"code": programs[program.ID-1].Code,
		}
		programList = append(programList, programData)
	}
	jsonData, err := json.Marshal(programList)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "err.tmpl", gin.H{})
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8005/moss", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Print(err)
		c.HTML(http.StatusInternalServerError, "err.tmpl", gin.H{})

	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		c.HTML(http.StatusInternalServerError, "err.tmpl", gin.H{})

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		c.HTML(http.StatusInternalServerError, "err.tmpl", gin.H{})
	}
	fmt.Println(string(body))
	var programsName = ""
	for _, program := range selectedPrograms {
		programsName += programs[program.ID-1].Name + " "
	}
	status := mossStatus.MossStatus{Status: "Выполнено", Link: string(body), Programs: programsName}
	h.services.StatusRep.Create(context.Background(), status)
	c.JSON(http.StatusOK, gin.H{"message": "Запрос успешно обработан"})
}

func (h *Handler) AntiplagiatStatus(c *gin.Context) {
	statuses, err := h.services.StatusRep.FindAll(context.Background())
	if err != nil {
		log.Print(err)
		c.HTML(http.StatusInternalServerError, "err.tmpl", gin.H{})
		return
	}

	squares := make([]SquaresStatus, 0)
	for _, p := range statuses {
		square := SquaresStatus{
			Date:   p.Date,
			Status: p.Status,
			Link:   p.Link,
		}
		squares = append(squares, square)
	}

	c.HTML(http.StatusOK, "Anti-plag.tmpl", gin.H{
		"blocks": squares,
	})
}
