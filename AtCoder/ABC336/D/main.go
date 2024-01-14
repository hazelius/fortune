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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	for i, a := range as {
		maxa := min(a, i+1)
		if i > 0 {
			maxa = min(maxa, as[i-1]+1)
		}
		as[i] = maxa
	}

	for i := n - 1; i >= 0; i-- {
		maxa := min(as[i], n-i)
		if i < n-1 {
			maxa = min(maxa, as[i+1]+1)
		}
		as[i] = maxa
	}

	ans := 0
	for _, a := range as {
		ans = max(ans, a)
	}

	fmt.Fprint(out, ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
