// L - Deque
// https://atcoder.jp/contests/dp/tasks/dp_l

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var dp [][]int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, as []int) int {
	dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	return f(0, n-1, as)
}

func f(s, e int, as []int) int {
	if e == s {
		dp[s][e] = as[s]
		return dp[s][e]
	}

	p1 := dp[s+1][e]
	if p1 == 0 {
		p1 = f(s+1, e, as)
	}

	p2 := dp[s][e-1]
	if p2 == 0 {
		p2 = f(s, e-1, as)
	}

	dp[s][e] = max(as[s]-p1, as[e]-p2)
	return dp[s][e]
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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, as))
}
