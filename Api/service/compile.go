package service

import (
	"Mehmat/structs"
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
)

type CompileService struct {
}

func NewCompileService() *CompileService {
	return &CompileService{}
}

func (s *CompileService) RunProgram(program structs.Program) (out string, err error) {

	//Chose a format of file which we will use
	var format string
	if program.Language == "c++" {
		format = "cpp"
	} else if program.Language == "go" {
		format = "go"
	} else {
		return "", errors.New("unknown language")
	}

	//Create a temporary file with code
	codeFile, err := os.CreateTemp("./Temp", "code.*."+format)
	if err != nil {
		return "", err
	}
	defer os.Remove(codeFile.Name())
	_, err = codeFile.WriteString(program.Code)
	if err != nil {
		return "", err
	}
	err = codeFile.Close()
	if err != nil {
		return "", err
	}

	//compile file
	var compiler string
	var flags []string
	if format == "cpp" {
		compiler = "g++"
		flags = []string{"-o", "program", codeFile.Name()}
	} else if format == "go" {
		compiler = "go"
		flags = []string{"build", "-o", "program", codeFile.Name()}
	}

	cmd := exec.Command(compiler, flags...)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	cmd = exec.Command("./program")
	cmd.Stdin = bytes.NewBufferString(program.StdIn)
	output, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	cmd.Start()
	outputBytes, _ := io.ReadAll(output)
	cmd.Wait()

	err = exec.Command("rm", "program").Run()
	if err != nil {
		return "", err
	}

	return string(outputBytes), nil
}
