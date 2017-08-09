package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	log.Printf("Start")
	flag.Parse()
	input_file := flag.Arg(0)
	output_file := flag.Arg(1)
	Read(input_file, output_file)

	dbconf, err := GetConf()
	if err != nil {
		log.Printf(err.Error())
	}
	fmt.Println(dbconf["1"])

	Connect()
  ssh_test()
	log.Printf("Finish !")
}
