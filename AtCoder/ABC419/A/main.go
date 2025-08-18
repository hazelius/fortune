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

	s := readString()
	switch s {
	case "red":
		fmt.Fprint(out, "SSS")
	case "blue":
		fmt.Fprint(out, "FFF")
	case "green":
		fmt.Fprint(out, "MMM")
	default:
		fmt.Fprint(out, "Unknown")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
