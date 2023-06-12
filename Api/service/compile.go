package service

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type CompileService struct {
}

func NewCompileService() *CompileService {
	return &CompileService{}
}

const dir = "./Temp"

func (s *CompileService) RunProgram(program []byte) (out string) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:1212/compile", bytes.NewBuffer(program))
	if err != nil {
		log.Print(err)
		return "error: compiler are not enable\n"
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return "error: Server Error\n"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return "error: Server Error\n"
	}
	return string(body)
}
