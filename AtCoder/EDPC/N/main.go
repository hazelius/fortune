// N - Slimes
// https://atcoder.jp/contests/dp/tasks/dp_n
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var dp [][]int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, as []int) int {
	dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	ans := f(0, n-1, as)
	return ans
}

func f(l, r int, as []int) int {
	if l == r {
		return 0
	}
	if dp[l][r] > 0 {
		return dp[l][r]
	}

	var cost = 0
	for i := l; i < r; i++ {
		lmin := dp[l][i]
		if lmin == 0 {
			lmin = f(l, i, as)
		}

		rmin := dp[i+1][r]
		if rmin == 0 {
			rmin = f(i+1, r, as)
		}

		sum := lmin + rmin
		if cost == 0 {
			cost = sum
		} else {
			cost = min(cost, sum)
		}
	}

	dp[l][r] = cost
	for i := l; i <= r; i++ {
		dp[l][r] += as[i]
	}
	return dp[l][r]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, as))
}
