package utils

import (
	"os"
	"os/exec"
)

func ClearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
