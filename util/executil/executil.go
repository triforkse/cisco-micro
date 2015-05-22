package executil

import (
	"os"
	"os/exec"
)

func Command(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
