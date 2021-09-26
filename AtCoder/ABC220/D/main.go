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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 10)
	}

	for i, a := range as {
		if i == 0 {
			dp[i][a]++
			continue
		}
		for preA, cnt := range dp[i-1] {
			if cnt == 0 {
				continue
			}
			t := (preA + a) % 10
			dp[i][t] = (dp[i][t] + cnt) % mod

			t = (preA * a) % 10
			dp[i][t] = (dp[i][t] + cnt) % mod
		}
	}

	ans := fmt.Sprintf("%v", dp[n-1])
	return ans[1 : len(ans)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
