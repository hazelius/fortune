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
	st := []string{readString(), readString()}
	sw := make([]int, n+1)
	for i := 0; i < m; i++ {
		l, r := readInt()-1, readInt()
		sw[l] = 1 - sw[l]
		sw[r] = 1 - sw[r]
	}

	idx := 0
	for i := 0; i < n; i++ {
		if sw[i] == 1 {
			idx = 1 - idx
		}
		fmt.Fprint(out, st[idx][i:i+1])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
