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

	readInt()
	s := readString()
	oneCnt := 0
	for _, c := range s {
		if c == '1' {
			oneCnt++
		}
	}

	idx := 0
	ans := 0
	for _, c := range s {
		if c == '1' {
			idx++
		} else {
			ans += min(idx, oneCnt-idx)
		}
	}

	fmt.Fprint(out, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
