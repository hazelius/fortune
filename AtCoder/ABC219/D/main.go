// D - Strange Lunchbox
// https://atcoder.jp/contests/abc219/tasks/abc219_d
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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	x, y := readInt(), readInt()

	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 301)
		for j := range dp[i] {
			dp[i][j] = make([]int, 301)
			for k := range dp[i][j] {
				dp[i][j][k] = n + 1
			}
		}
	}

	dp[0][0][0] = 0
	for i := 1; i <= n; i++ {
		a, b := readInt(), readInt()
		for ia := range dp[i] {
			for ib, v := range dp[i][ia] {
				dp[i][min(ia+a, x)][min(ib+b, y)] = min(dp[i][min(ia+a, x)][min(ib+b, y)], dp[i-1][ia][ib]+1)
				dp[i][ia][ib] = min(v, dp[i-1][ia][ib])
			}
		}
	}

	ans := dp[n][x][y]
	if ans > n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
