package utils

import (
	"fmt"
	"os/exec"
)

func Extract(path string, to string) {

	cmd := exec.Command("tar", "-xvf", path, "-C", to)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
