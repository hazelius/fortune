package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	for i, c := range s {
		if i == 0 {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			if !strings.Contains(t, s[i-1:i]) {
				fmt.Fprint(out, "No")
				return
			}
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
