package service

import (
	"Mehmat/structs"
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type CompileService struct {
}

func NewCompileService() *CompileService {
	return &CompileService{}
}

const dir = "./Temp"

func (s *CompileService) RunProgram(program structs.Program) (out string, err error) {
	lang := strings.ToLower(program.Language)
	format, compiler, flags := "", "", []string{}

	//Chose a format of file which will used
	switch lang {
	case "c++":
		format = "cpp"
		compiler = "g++"
		flags = []string{"-o", "program"}
	case "go":
		format = "go"
		compiler = "go"
		flags = []string{"build", "-o", "program"}
	default:
		return "", errors.New("unknown language")
	}

	//Create temp file
	codeFile, err := os.CreateTemp(dir, "code.*."+format)
	if err != nil {
		return "", err
	}
	defer os.Remove(codeFile.Name())

	if _, err = codeFile.WriteString(program.Code); err != nil {
		return "", err
	}
	if err = codeFile.Close(); err != nil {
		return "", err
	}

	flags = append(flags, codeFile.Name())

	//Build Program
	cmd := exec.Command(compiler, flags...)
	err = cmd.Run()
	if err != nil {
		log.Printf("Error building program: %v", err)
		return "", err
	}

	//Run Program
	cmd = exec.Command("./program")

	cmd.Stdin = bytes.NewBufferString(program.StdIn)
	output, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	outputBytes, err := io.ReadAll(output)
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		log.Printf("Error running program: %v", err)
		return "", err
	}

	return string(outputBytes), nil
}
