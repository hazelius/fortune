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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s, t := readString(), readString()
	if s == t {
		fmt.Fprint(out, 0)
		return
	}

	for i := 0; i < 100; i++ {
		if len(s) <= i || len(t) <= i {
			fmt.Fprint(out, i+1)
			return
		}
		if s[i] != t[i] {
			fmt.Fprint(out, i+1)
			return
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
