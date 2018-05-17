package main

import (
	"fmt"
	"log"
)

func main() {
	log.Print("Start")
	fmt.Printf("hello, world\n")
	name := fullname("John", "Smith")
	fmt.Print("hello " + name)
}

func fullname(firstname, lastname string) string {
	return fmt.Sprintf("%v %v", firstname, lastname)
}
