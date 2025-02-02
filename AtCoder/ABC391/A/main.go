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
	d := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}

	for i, v := range d {
		if s == v {
			fmt.Fprint(out, d[(i+4)%8])
			return
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
