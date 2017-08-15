package main

import (
	"flag"
	"log"
)

func main() {
	var err error
	log.Printf("Start")

	flag.Parse()
	input_file := flag.Arg(0)
	output_file := flag.Arg(1)
	Read(input_file, output_file)

	Connect()
  conn, err := ssh_test()
  if err != nil {
    log.Println(err)
  }
  defer conn.Close()
  ftp_test(conn)

	err = Download("test_cydas", "test.txt")
	if err != nil {
	    log.Printf("failed to download file, %v", err)
	}
	log.Printf("Finish !")
}
