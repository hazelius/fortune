package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	if strings.HasSuffix(s, "er") {
		return "er"
	}
	if strings.HasSuffix(s, "ist") {
		return "ist"
	}
	return ""
}

func main() {
	fmt.Println(run(os.Stdin))
}
