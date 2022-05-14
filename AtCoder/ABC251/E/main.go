package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	dp := make([][]int, n+1)
	dp[0] = []int{0, 1e9}
	for i, a := range as {
		dp[i+1] = []int{0, 0}
		dp[i+1][1] = min(dp[i][0], dp[i][1]) + a
		dp[i+1][0] = dp[i][1]
	}
	ans1 := dp[n][0]

	dp = make([][]int, n+1)
	dp[0] = []int{1e9, 0}
	for i, a := range as {
		dp[i+1] = []int{0, 0}
		dp[i+1][1] = min(dp[i][0], dp[i][1]) + a
		dp[i+1][0] = dp[i][1]
	}
	ans2 := dp[n][1]

	return min(ans1, ans2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
