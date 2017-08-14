package main

import (
	"flag"
	"log"
)

func main() {
	log.Printf("Start")
	flag.Parse()
	input_file := flag.Arg(0)
	output_file := flag.Arg(1)
	Read(input_file, output_file)


	InitConf()

	Connect()
  conn, err := ssh_test()
  if err != nil {
    log.Println(err)
  }
  defer conn.Close()
  ftp_test(conn)

	log.Printf("Finish !")
}
