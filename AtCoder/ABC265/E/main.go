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

type point struct {
	x int
	y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	dirs := make([][]int, 3)
	for i := range dirs {
		dirs[i] = []int{readInt(), readInt()}
	}
	pointmap := make(map[point]bool)
	for i := 0; i < m; i++ {
		p := point{readInt(), readInt()}
		pointmap[p] = true
	}

	dp := make([]map[point]int, n+1)
	dp[0] = make(map[point]int)
	dp[0][point{0, 0}] = 1

	for i := 1; i <= n; i++ {
		dp[i] = make(map[point]int)
		for p, cnt := range dp[i-1] {
			for _, d := range dirs {
				next := point{p.x + d[0], p.y + d[1]}
				if pointmap[next] {
					continue
				}
				nextcnt, ok := dp[i][next]
				if !ok {
					dp[i][next] = cnt
				} else {
					dp[i][next] = (nextcnt + cnt) % mod
				}
			}
		}
	}

	ans := 0
	for _, cnt := range dp[n] {
		ans = (ans + cnt) % mod
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
