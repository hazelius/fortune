package main

import (
	"fmt"
	"os/exec"
)

func main() {
	sudoCommand()
	dockerCommand()
}

func sudoCommand() {
	// コメント
	cmd := "/bin/sh"
	opt := []string{"-c", "sudo ls"}
	out, err := exec.Command(cmd, opt...).CombinedOutput()
	if err != nil {
		fmt.Printf("エラー:%v\n", err)
	}
	fmt.Println(string(out))
}

func dockerCommand() {
	fmt.Println("======= 01 Run")
	out, err := exec.Command("cmd", "/C", "docker", "ps").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
