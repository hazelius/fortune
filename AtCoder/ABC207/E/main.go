package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const mod = 1e9 + 7

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
	sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + readInt()
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	for i, sum := range sums {
		mem := make([]int, i+1)
		mem[sum%(i+1)] = dp[i][i]

		for j := i + 1; j <= n; j++ {
			k := sums[j] % (i + 1)
			dp[i+1][j] = mem[k]
			mem[k] = (mem[k] + dp[i][j]) % mod
		}
	}

	ans := 0
	for _, v := range dp {
		ans = (ans + v[n]) % mod
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
