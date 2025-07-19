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
	idx := 0
	ans := make([]int, 2)
	for i, c := range s {
		if c == '#' {
			ans[idx] = i + 1
			idx++
			if idx > 1 {
				fmt.Fprintf(out, "%v,%v\n", ans[0], ans[1])
				idx = 0
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
