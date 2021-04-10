// https://atcoder.jp/contests/dp/tasks/dp_e

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

func run(n, w int, wvs [][]int) int {
	dp := make([][]int, n+1)
	dp[0] = make([]int, 100001)
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = w + 1
	}

	for i, wv := range wvs {
		dp[i+1] = make([]int, 100001)
		for v1, w1 := range dp[i] {
			if v1-wv[1] >= 0 {
				dp[i+1][v1] = min(w1, dp[i][v1-wv[1]]+wv[0])
			} else {
				dp[i+1][v1] = w1
			}
		}
	}

	for i := 100000; i >= 0; i-- {
		if dp[n][i] <= w {
			return i
		}
		// fmt.Println(dp[n][i])
	}
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n, w := readInt(), readInt()
	wvs := make([][]int, n)
	for i := 0; i < n; i++ {
		wvs[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(n, w, wvs))
}
