package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("======= 01 Run")
	out, err := exec.Command("cmd", "/C", "docker", "ps").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
