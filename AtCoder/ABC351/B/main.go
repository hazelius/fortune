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
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	for i := 0; i < n; i++ {
		s := readString()
		a := ss[i]
		for j := range s {
			if a[j:j+1] != s[j:j+1] {
				fmt.Fprint(out, i+1, j+1)
				return
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
