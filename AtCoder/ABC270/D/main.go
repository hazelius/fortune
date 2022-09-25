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

	n, k := readInt(), readInt()
	as := make([]int, k)
	for i := range as {
		as[i] = readInt()
	}

	dp := make([][]int, n+1)

	for i := range dp {
		if i == 0 {
			dp[i] = []int{0, 0}
			continue
		}

		dp[i] = []int{0, n}
		for _, a := range as {
			if a > i {
				continue
			}

			dp[i][0] = max(dp[i][0], dp[i-a][1]+a)
			dp[i][1] = min(dp[i][1], dp[i-a][0])
		}
	}

	fmt.Fprint(out, dp[n][0])
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
