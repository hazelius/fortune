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

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i, a := range as {
		if i == 0 {
			dp[i][1] = a
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+a+a)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+a)
	}

	fmt.Fprint(out, max(dp[n-1][0], dp[n-1][1]))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
