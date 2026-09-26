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

	c := readString()
	switch c {
	case "B":
		fmt.Fprint(out, "Y")
	case "Y":
		fmt.Fprint(out, "R")
	case "R":
		fmt.Fprint(out, "B")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
