package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	s := readString()

	for i := 1; i < n; i++ {
		fmt.Fprintln(out, f(n, i, s))
	}
}

func f(n, i int, s string) int {
	l := 0
	for ; l+i < n; l++ {
		if s[l] == s[l+i] {
			return l
		}
	}
	return l
}

func main() {
	run(os.Stdin, os.Stdout)
}
