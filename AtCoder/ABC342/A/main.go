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

	s := readString()
	for i := range s {
		if i == 0 {
			continue
		}
		if s[i] == s[i-1] {
			continue
		}
		var c byte
		if i > 1 {
			c = s[i-2]
		} else {
			c = s[i+1]
		}
		if s[i] == c {
			fmt.Fprint(out, i)
		} else {
			fmt.Fprint(out, i+1)
		}
		return
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
