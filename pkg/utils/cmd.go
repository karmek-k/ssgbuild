package utils

import (
	"os/exec"
	"strings"
)

// StringToCmd converts a command string to exec.Cmd
func StringToCmd(cmdString string) *exec.Cmd {
	tokens := strings.Split(cmdString, " ")

	if len(tokens) == 1 {
		return exec.Command(tokens[0])
	}

	return exec.Command(tokens[0], tokens[1:]...)
}
