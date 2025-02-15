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

	s1, s2 := readString(), readString()
	ans := 1
	if s1 == "fine" {
		if s2 == "fine" {
			ans = 4
		} else {
			ans = 3
		}
	} else {
		if s2 == "fine" {
			ans = 2
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
