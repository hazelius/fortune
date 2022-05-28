package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var mod = 998244353

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m, k := readInt(), readInt(), readInt()
	dp := make([][]int, 2)
	dp[0] = make([]int, m)
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 0; i < n-1; i++ {
		idx := i % 2
		nextIdx := 1 - idx
		dp[nextIdx] = make([]int, m)

		dps := make([]int, m+1)
		for j := 0; j < m; j++ {
			dps[j+1] = dps[j] + dp[idx][j]
		}

		for j := range dp[nextIdx] {
			maxa := j + 1 - k
			if maxa > 0 {
				dp[nextIdx][j] = (dps[maxa] - dps[0]) % mod
			}
			mina := j + k
			if mina < m {
				dp[nextIdx][j] = (dp[nextIdx][j] + dps[m] - dps[mina]) % mod
			}
		}
	}

	ans := 0
	for _, v := range dp[1-n%2] {
		ans = (ans + v) % mod
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
