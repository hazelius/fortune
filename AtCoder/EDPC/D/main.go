// https://atcoder.jp/contests/dp/tasks/dp_d

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
	dp[0] = make([]int, w+1)
	for i, wv := range wvs {
		dp[i+1] = make([]int, w+1)
		for wi := 0; wi <= w; wi++ {
			if wi-wv[0] >= 0 {
				dp[i+1][wi] = max(dp[i][wi], dp[i][wi-wv[0]]+wv[1])
			} else {
				dp[i+1][wi] = dp[i][wi]

			}
		}
	}
	// fmt.Println(dp[n-1])
	return dp[n][w]
}

func max(a, b int) int {
	if a < b {
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
