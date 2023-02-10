package lib

import (
	"os/exec"
)

func RunApp(pathToBin string, stdin string) (out string, err error) {

	//Check path to bin
	path, err := exec.LookPath(pathToBin)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(path)

	//Write data with pipeIn
	pipeIn, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	pipeIn.Write([]byte(stdin))
	pipeIn.Close()

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
