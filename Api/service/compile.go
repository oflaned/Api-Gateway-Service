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

	//Chose a format of file which will used
	var format string
	switch program.Language {
	case "c++":
		format = "cpp"
	case "go":
		format = "go"
	default:
		return "", errors.New("unknown language")
	}

	//Create a temporary folder and file with code
	tempDir, err := os.MkdirTemp("", "Temp")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tempDir)

	codeFile, err := os.CreateTemp(tempDir, "code.*."+format)
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

	//Compile file
	var compiler string
	var flags []string
	switch format {
	case "cpp":
		compiler = "g++"
		flags = []string{"-o", "program", codeFile.Name()}
	case "go":
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

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	outputBytes, err := io.ReadAll(output)
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	err = exec.Command("rm", "program").Run()
	if err != nil {
		return "", err
	}

	return string(outputBytes), nil
}
