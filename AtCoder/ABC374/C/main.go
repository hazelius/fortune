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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	m := 0
	for i := range as {
		a := readInt()
		as[i] = a
		m += a
	}

	dp := make([]map[int]bool, n+1)
	dp[0] = make(map[int]bool)
	dp[0][0] = true
	for i, a := range as {
		dp[i+1] = make(map[int]bool)
		for k := range dp[i] {
			dp[i+1][k] = true
			dp[i+1][k+a] = true
		}
	}
	ans := m
	m = (m + 1) / 2
	for k := range dp[n] {
		if k >= m && ans > k {
			ans = k
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
