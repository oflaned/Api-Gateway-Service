package lib

import (
	"log"
	"os/exec"
)

func BuildGoApp(pathToProj string) (result bool, err error) {
	cmd := exec.Command("go", "build")
	cmd.Dir = pathToProj
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	log.Println("Bin has been compiled successfully: ", pathToProj, string(output))
	return true, nil
}
