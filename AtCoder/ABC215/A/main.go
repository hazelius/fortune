package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	s := readString()
	if s == "Hello,World!" {
		return "AC"
	}
	return "WA"
}

func main() {
	fmt.Println(run(os.Stdin))
}
