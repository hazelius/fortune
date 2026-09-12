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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	s := readString()
	for i, c := range s {
		if i > 0 {
			fmt.Fprint(out, "o")
		}
		fmt.Fprint(out, string(c))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
