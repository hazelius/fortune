// O - Matching
// https://atcoder.jp/contests/dp/tasks/dp_o
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const mod = 1000000007

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, as [][]int) int {
	dp := make([][]int, n+1)
	dp[0] = make([]int, 1<<n)
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		dp[i+1] = make([]int, 1<<n)
		for s := 0; s < 1<<n; s++ {
			for j := 0; j < n; j++ {
				if (s>>j)&1 == 1 && as[i][j] == 1 {
					dp[i+1][s] = (dp[i+1][s] + dp[i][s^(1<<j)]) % mod
				}
			}
		}
	}

	// fmt.Println(dp)
	return dp[n][(1<<n)-1]
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	as := make([][]int, n)
	for i := range as {
		as[i] = make([]int, n)
		for j := range as[i] {
			as[i][j] = readInt()
		}
	}
	fmt.Println(run(n, as))
}
