package main 

import (
	"fmt"
	"golang.org/x/crypto/ssh"
//	"golang.org/x/crypto/ssh/agent"
	"bytes"
	"io"
	"os"
)

func main() {
	config := &ssh.ClientConfig{
		User: "xxxx",
		Auth: []ssh.AuthMethod{
			ssh.Password("xxxx"),
		},
	}
	modes := ssh.TerminalModes{
		ssh.ECHO: 0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	client, err := ssh.Dial("tcp","xxxx:22", config)
	if err != nil {
		panic("Failed to Dial: " + err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
		panic(err.Error())
	}
	stdin, err := session.StdinPipe()
if err != nil {
	panic( err)
}
go io.Copy(stdin, os.Stdin)

stdout, err := session.StdoutPipe()
if err != nil {
	panic(err)
}
go io.Copy(os.Stdout, stdout)

stderr, err := session.StderrPipe()
if err != nil {
	panic(err)
}
go io.Copy(os.Stderr, stderr)
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/bin/bash;ls -lah"); err != nil {
		panic("Failed to run:" + err.Error())
	}
	
	fmt.Println(b.String())
}

