package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	n, m := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			dp[i][j] = math.MinInt64
		}
	}

	for i, a := range as {
		for j := range dp[i] {
			if j == 0 {
				continue
			}
			if i+1 < j {
				break
			}

			if i == 0 {
				dp[i][j] = a
				continue
			}

			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+a*j)
		}
	}

	fmt.Fprint(out, dp[n-1][m])
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
