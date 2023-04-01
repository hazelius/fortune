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

	for _, a := range "87654321" {
		s := readString()
		for j, b := range "abcdefgh" {
			if s[j] == '*' {
				fmt.Fprintf(out, "%v%v", string(b), string(a))
				return
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
