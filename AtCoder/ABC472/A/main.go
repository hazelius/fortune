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
	for _, c := range s {
		if c == 'A' {
			fmt.Fprint(out, "A")
		} else {
			fmt.Fprint(out, ".")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
