package main

import (
  "flag"
  "log"

  "./csv"
  "./mysql"
)

func main() {
  log.Printf("Start")
  flag.Parse()
  input_file := flag.Arg(0)
  output_file := flag.Arg(1)
  csv.Read(input_file, output_file)

  db.Connect()
  log.Printf("Finish !")
}
