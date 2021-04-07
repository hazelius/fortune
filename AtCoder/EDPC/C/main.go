// https://atcoder.jp/contests/dp/tasks/dp_c

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

func run(n int, abc [][]int) int {
	dp := make([][]int, n+1)
	dp[0] = abc[0]
	for i := 1; i < n; i++ {
		dp[i] = make([]int, len(abc[i]))
		preMax := []int{
			max(dp[i-1][1], dp[i-1][2]),
			max(dp[i-1][0], dp[i-1][2]),
			max(dp[i-1][0], dp[i-1][1]),
		}
		for j := 0; j < len(abc[i]); j++ {
			dp[i][j] = abc[i][j] + preMax[j]
		}
	}
	r := dp[n-1]
	return max(max(r[0], r[1]), r[2])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	abc := make([][]int, n)
	for i := 0; i < n; i++ {
		abc[i] = []int{readInt(), readInt(), readInt()}
	}
	fmt.Println(run(n, abc))
}
