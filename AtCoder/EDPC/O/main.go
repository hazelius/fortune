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
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 1<<n)
	}
	dp[0][0] = 1
	for s := 1; s < 1<<n; s++ {
		i := bitcount(s) - 1
		for j := 0; j < n; j++ {
			if (s>>j)&1 == 0 || as[i][j] == 0 {
				continue
			}
			dp[i+1][s] = (dp[i+1][s] + dp[i][s^(1<<j)]) % mod
		}
	}

	// fmt.Println(dp)
	return dp[n][(1<<n)-1]
}

func bitcount(s int) int {
	cnt := 0
	for i := 0; i < 21; i++ {
		if s>>i&1 == 1 {
			cnt++
		}
	}
	return cnt
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
