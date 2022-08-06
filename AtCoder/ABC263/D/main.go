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

	n, l, r := readInt(), readInt(), readInt()
	as := make([]int, n)
	sum := 0
	for i := range as {
		a := readInt()
		as[i] = a
		sum += a
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	for i := 0; i < n; i++ {
		x := l - as[i]
		if i == 0 {
			dp[i][1] = x
			dp[i][0] = min(dp[i][0], x)
		} else {
			dp[i][1] = dp[i-1][1] + x
			dp[i][0] = min(dp[i-1][0], min(dp[i-1][1], dp[i][1]))
		}
	}

	for i := n - 1; i >= 0; i-- {
		y := r - as[i]
		if i == n-1 {
			dp[i][2] = y
			dp[i][0] = min(dp[i][0], min(dp[i][2], y))
		} else {
			dp[i][2] = dp[i+1][2] + y
			cmin := min(dp[i+1][2], dp[i][2])
			if i > 0 {
				cmin += dp[i-1][0]
			}
			dp[i][0] = min(dp[i][0], cmin)
		}
	}

	ans := 0
	for _, d := range dp {
		ans = min(ans, d[0])
	}

	fmt.Fprint(out, ans+sum)
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
