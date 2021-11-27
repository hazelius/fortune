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

	s1, s2 := readString(), readString()
	if s1[0] == '.' && s2[1] == '.' {
		return "No"
	}
	if s1[1] == '.' && s2[0] == '.' {
		return "No"
	}
	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
