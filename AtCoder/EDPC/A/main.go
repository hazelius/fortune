// https://atcoder.jp/contests/dp/tasks/dp_a
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

func run(n int, h []int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = abs(h[1] - h[0])
	for i := 2; i < n; i++ {
		dp[i] = min(
			dp[i-1]+abs(h[i-1]-h[i]),
			dp[i-2]+abs(h[i-2]-h[i]),
		)
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = readInt()
	}
	fmt.Println(run(n, h))
}
