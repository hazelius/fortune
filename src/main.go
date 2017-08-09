package main

import (
	"flag"
	"fmt"
	"log"
  "golang.org/x/crypto/ssh"
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
  var conn *ssh.Client
  conn, err = ssh_test()
  if err != nil {
    log.Println(err)
  }
  defer conn.Close()
  ftp_test(conn)

	log.Printf("Finish !")
}
