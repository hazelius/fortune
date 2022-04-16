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
	if m > k {
		m = k
	}

	dp := make([][]int, n+1)
	dp[0] = make([]int, k+1)
	for i := 1; i <= m; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		dp[i] = make([]int, k+1)

		for j := 1; j <= k; j++ {
			v := dp[i-1][j]
			if v == 0 {
				continue
			}
			for p := 1; p <= m; p++ {
				if j+p <= k {
					dp[i][j+p] = (dp[i][j+p] + v) % mod
				}
			}
		}
	}

	ans := 0
	for _, v := range dp[n-1] {
		ans = (ans + v) % mod
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
