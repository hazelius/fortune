// https://atcoder.jp/contests/dp/tasks/dp_g
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

func run(n, m int, xys [][]int) int {
	links := make(map[int][]int)
	for _, xy := range xys {
		x, y := xy[0], xy[1]
		if v, ok := links[x]; ok {
			links[x] = append(v, y)
		} else {
			links[x] = []int{y}
		}
	}

	dp := make(map[int]int)
	ans := 0
	for x := 1; x <= n; x++ {
		ans = max(ans, f(x, dp, links))
	}
	return ans
}

func f(x int, dp map[int]int, links map[int][]int) int {
	if v, ok := dp[x]; ok {
		return v
	}
	dp[x] = 0

	l, _ := links[x]
	for _, y := range l {
		dp[x] = max(dp[x], f(y, dp, links)+1)
	}
	return dp[x]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := readInt(), readInt()
	xys := make([][]int, m)
	for i := 0; i < m; i++ {
		xys[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(n, m, xys))
}
