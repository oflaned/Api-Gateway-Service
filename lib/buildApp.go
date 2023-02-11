package lib

import (
	"log"
	"os/exec"
)

func BuildGoApp(pathToProj string, folder string) (pathToBin string, err error) {
	cmd := exec.Command("go", "build")
	cmd.Dir = pathToProj + folder
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	log.Println("Bin has been compiled successfully: ", pathToProj+folder, string(output))
	return folder, nil
}
