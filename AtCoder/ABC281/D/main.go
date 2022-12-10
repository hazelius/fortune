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

	n, k, d := readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, d)
			for idx := range dp[i][j] {
				dp[i][j][idx] = -1
			}
		}
	}
	dp[0][0][0] = 0

	for i := 0; i < n; i++ {
		a := as[i]
		for j := 0; j <= k; j++ {
			for am := range dp[i][j] {
				if dp[i][j][am] == -1 {
					continue
				}
				dp[i+1][j][am] = max(dp[i+1][j][am], dp[i][j][am])
				if j != k {
					newam := (am + a) % d
					dp[i+1][j+1][newam] = max(dp[i+1][j+1][newam], dp[i][j][am]+a)
				}
			}
		}
	}

	fmt.Fprint(out, dp[n][k][0])
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
