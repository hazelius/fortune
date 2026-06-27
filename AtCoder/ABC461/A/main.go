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
	w, e := 0, 0
	for _, c := range s {
		if c == 'E' {
			e++
		} else {
			w++
		}
	}
	if w > e {
		fmt.Fprint(out, "West")
	} else {
		fmt.Fprint(out, "East")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
