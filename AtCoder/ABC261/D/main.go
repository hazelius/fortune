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

	n, m := readInt(), readInt()
	xs := make([]int, n)
	for i := range xs {
		xs[i] = readInt()
	}

	cymap := make(map[int]int)
	for i := 0; i < m; i++ {
		c, y := readInt(), readInt()
		cymap[c-1] = y
	}

	dp := make([][]int, n+1)
	dp[0] = make([]int, 1)
	for i := 0; i < n; i++ {
		dp[i+1] = make([]int, i+2)
		x := xs[i]
		for c, s := range dp[i] {
			snew := s + x
			y, ok := cymap[c]
			if ok {
				snew += y
			}
			if c+1 < len(dp[i]) {
				snew = max(snew, dp[i][c+1])
			}
			dp[i+1][c+1] = snew

			snew = max(s, dp[i][0])
			dp[i+1][0] = max(snew, dp[i+1][0])
		}
	}

	ans := 0
	for _, v := range dp[n] {
		ans = max(ans, v)
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
