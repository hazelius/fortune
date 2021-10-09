package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var mod = 998244353

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
	bs := make([]int, n)
	for i := range bs {
		bs[i] = readInt()
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, bs[n-1]+1)
	}

	for a := as[0]; a <= bs[0]; a++ {
		dp[0][a] = 1
	}

	for i := 1; i < n; i++ {
		a := as[i]
		for j := 0; j <= as[i]; j++ {
			dp[i][a] = (dp[i][a] + dp[i-1][j]) % mod
		}
		for j := a + 1; j <= bs[i]; j++ {
			dp[i][j] = (dp[i][j-1] + dp[i-1][j]) % mod
		}
	}

	ans := 0
	for i := as[n-1]; i <= bs[n-1]; i++ {
		ans = (ans + dp[n-1][i]) % mod
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
