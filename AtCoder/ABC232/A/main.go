package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	s := readString()
	a, _ := strconv.Atoi(s[:1])
	b, _ := strconv.Atoi(s[2:])
	return a * b
}

func main() {
	fmt.Println(run(os.Stdin))
}
