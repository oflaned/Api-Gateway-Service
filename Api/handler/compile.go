package handler

import (
	"Mehmat/lib"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

var bannedFunctions = map[string]bool{
	"new":     true,
	"delete":  true,
	"malloc":  true,
	"free":    true,
	"calloc":  true,
	"realloc": true,
	"strcpy":  true,
	"strcat":  true,
	"sprintf": true,
	"fopen":   true,
	"fread":   true,
	"fwrite":  true,
	"fclose":  true,
	"fork":    true,
	"exec":    true,
	"system":  true,
}

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

	if lib.ContainsBannedFunctions(code, bannedFunctions) {
		log.Print("Used of banned functions is not allowed")
		c.String(http.StatusBadRequest, "Used banned functions \n")
		return
	}

	// Данные для отправки
	data := map[string]string{
		"code":  code,
		"lang":  lang,
		"input": input,
	}

	// Кодируем данные в JSON
	program, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	// Создаем клиент
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:1212/compile", bytes.NewBuffer(program))
	if err != nil {
		panic(err)
	}

	// Добавляем заголовки
	req.Header.Add("Content-Type", "application/json")

	// Отправляем запрос и получаем ответ
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	c.String(http.StatusOK, string(body))
}

func (h *Handler) Program(c *gin.Context) {
	c.HTML(http.StatusOK, "program.tmpl", gin.H{})
}
