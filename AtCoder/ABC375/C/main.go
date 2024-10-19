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
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([][]byte, n)
	for i := range as {
		as[i] = []byte(readString())
	}

	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, n)
	}

	for i, a := range as {
		for j, c := range a {
			l := level(i, j, n)
			x, y := i, j
			for k := 0; k < l; k++ {
				x, y = rotate(x, y, n)
			}
			ans[x][y] = c
		}
	}

	for _, v := range ans {
		fmt.Fprintln(out, string(v))
	}
}

func level(i, j, n int) int {
	ret := min(i, j)
	ret = min(ret, n-1-i)
	ret = min(ret, n-1-j)
	return 1 + ret%4
}

func rotate(i, j, n int) (int, int) {
	return j, n - 1 - i
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
