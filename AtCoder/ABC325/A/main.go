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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	s, _ := readString(), readString()
	fmt.Fprintf(out, "%s san", s)
}

func main() {
	run(os.Stdin, os.Stdout)
}
