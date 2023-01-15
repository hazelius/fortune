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

	turnmap := make(map[int]int)

	for i := 1; i < n; i++ {
		turn := 0
		for j := 0; j < i; j++ {
			turn += as[min(j, i-j-1)]
		}
		turnmap[i] = turn
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	dp[1][0] = 0
	for i := 1; i < n; i++ {
		for j := 0; j <= n; j++ {
			if dp[i][j] < 0 {
				continue
			}
			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j])
			dp[i+1][0] = max(dp[i+1][0], dp[i][j]+turnmap[j])
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, dp[n][i]+turnmap[i])
	}

	fmt.Fprint(out, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
