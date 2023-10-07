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
	for i, v := range s {
		if i%2 == 1 && v != '0' {
			fmt.Fprint(out, "No")
			return
		}
	}
	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
