// https://atcoder.jp/contests/dp/tasks/dp_i
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

func readFloat() float64 {
	sc.Scan()
	i, _ := strconv.ParseFloat(sc.Text(), 64)
	return i
}

func run(n int, ps []float64) float64 {
	dp := make([][]float64, n+1)
	dp[0] = []float64{1}
	for i, p := range ps {
		idx := i + 1
		dp[idx] = make([]float64, idx+1)

		l := idx - n/2
		if l < 0 {
			l = 0
		}
		for j := l; j <= idx; j++ {
			if j == 0 {
				dp[idx][0] = dp[idx-1][0] * (1 - p)
			} else if j == idx {
				dp[idx][j] = dp[idx-1][j-1] * p
			} else {
				dp[idx][j] = dp[idx-1][j]*(1-p) + dp[idx-1][j-1]*p
			}
		}
	}
	var ans float64
	for i := n/2 + 1; i <= n; i++ {
		ans += dp[n][i]
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	ps := make([]float64, n)
	for i := range ps {
		ps[i] = readFloat()
	}
	fmt.Println(run(n, ps))
}
