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

type road struct {
	a int
	b int
	c int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m, k := readInt(), readInt(), readInt()
	roads := make([]road, m)
	for i := range roads {
		roads[i] = road{readInt() - 1, readInt() - 1, readInt()}
	}

	ks := make([]int, k)
	for i := range ks {
		ks[i] = readInt() - 1
	}

	dp := make([]int, n)
	for _, k := range ks {
		r := roads[k]
		if r.a == 0 {
			dp[r.b] = min(dp[r.b], r.c)
		} else if dp[r.a] > 0 {
			dp[r.b] = min(dp[r.b], r.c+dp[r.a])
		}
	}

	ans := dp[n-1]
	if ans == 0 {
		fmt.Fprint(out, -1)
		return
	}
	fmt.Fprint(out, ans)
}

func min(a, b int) int {
	if a == 0 {
		return b
	}
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
