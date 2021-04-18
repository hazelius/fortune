// M - Candies
// https://atcoder.jp/contests/dp/tasks/dp_m

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n, k int, as []int) int {
	mod := 1000000007
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			for a := 0; a <= as[i]; a++ {
				if j-a < 0 {
					continue
				}
				if i == 0 {
					if j == a {
						dp[i][j] += 1
					}
				} else {
					dp[i][j] += dp[i-1][j-a]
				}
				dp[i][j] %= mod
			}
		}
	}
	return dp[n-1][k]
}

func main() {
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, k, as))
}
