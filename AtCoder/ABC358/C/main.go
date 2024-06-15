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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	ss := make([]int, n)
	for i := range ss {
		s := readString()
		for j := range s {
			if s[j] == 'o' {
				ss[i] |= 1 << j
			}
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 1<<m)
		for j := range dp[i] {
			dp[i][j] = 1 << m
		}
	}

	for i, s := range ss {
		dp[i+1][s] = 1
		for j, v := range dp[i] {
			dp[i+1][j] = min(dp[i+1][j], dp[i][j])
			j |= s
			v++
			dp[i+1][j] = min(dp[i+1][j], v)
		}
	}

	fmt.Fprint(out, dp[n][1<<m-1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
