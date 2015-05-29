package executil

import (
	"os"
	"os/exec"
        "log"
)

func Command(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func CommandList(args ...string) *exec.Cmd {
        if len(args) == 0 {
                log.Fatal("what are you doing?")
        }
        command := args[0]
        cmd := exec.Command(command, args[1:]...)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        return cmd
}
