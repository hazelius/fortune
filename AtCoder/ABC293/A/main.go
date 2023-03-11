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
	ans := make([]byte, len(s))
	for i, c := range s {
		if i%2 == 0 {
			ans[i+1] = byte(c)
		} else {
			ans[i-1] = byte(c)
		}
	}

	fmt.Fprint(out, string(ans))
}

func main() {
	run(os.Stdin, os.Stdout)
}
