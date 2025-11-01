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

	n, m := readInt(), readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	smap := make(map[string]bool)
	for i := 0; i < n-m+1; i++ {
		for j := 0; j < n-m+1; j++ {
			str := ""
			for k := i; k < i+m; k++ {
				str += ss[k][j : j+m]
			}
			smap[str] = true
		}
	}

	fmt.Fprint(out, len(smap))
}

func main() {
	run(os.Stdin, os.Stdout)
}
