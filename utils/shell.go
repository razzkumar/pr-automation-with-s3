package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// Execute runs the process with the supplied environment.
func RunCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stderr = &stderr
	cmd.Stdout = &out

	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())

		return errors.New(err.Error())
	}

	return nil
}
