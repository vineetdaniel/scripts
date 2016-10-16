package main

import (
	"fmt"
	"os"
	"os/exec"
)

type SSHCommander struct {
	User string
	IP   string
}

func (s *SSHCommander) Command(cmd ...string) *exec.Cmd {
	arg := append(
		[]string{
			fmt.Sprintf("%s@%s", s.User, s.IP),
		},
		cmd...,
	)
	return exec.Command("ssh", arg...)
}

func main() {
	commander := SSHCommander{"root", "localhost"}

	cmd := []string{
		"ls",
		"-la",
	}
	fmt.Println(cmd)
	// am I doing this automation thing right?
	if err := commander.Command("ls -la"); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running SSH command: ", err)
		os.Exit(1)
	}
}
