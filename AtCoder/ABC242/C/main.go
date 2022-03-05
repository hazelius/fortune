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

	n := readInt()
	dp := make([][]int, n+1)
	dp[1] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		dp[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			dp[i][j] = (dp[i][j] + dp[i-1][j]) % mod
			if j > 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if j < 8 {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}

	ans := 0
	for _, v := range dp[n] {
		ans = (ans + v) % mod
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
