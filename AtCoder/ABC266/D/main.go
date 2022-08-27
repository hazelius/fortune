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
	maptxa := make(map[int][]int)
	maxt := 0
	for i := 0; i < n; i++ {
		t, x, a := readInt(), readInt(), readInt()
		maxt = max(maxt, t)
		maptxa[t] = []int{x, a}
	}

	dp := [][]int{
		{0, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1},
	}
	slot := 0
	pre := 2
	for time := 0; time <= maxt; time++ {
		v, ok := maptxa[time]
		if ok {
			x, a := v[0], v[1]
			if dp[slot][x] >= 0 {
				dp[slot][x] += a
			}
		}

		pre = slot
		slot = (slot + 1) % 3

		for pos, v := range dp[pre] {
			dp[slot][pos] = v
			if pos > 0 {
				dp[slot][pos] = max(dp[slot][pos], dp[pre][pos-1])
			}
			if pos < 4 {
				dp[slot][pos] = max(dp[slot][pos], dp[pre][pos+1])
			}
		}
	}

	ans := 0
	for _, v := range dp[pre] {
		ans = max(ans, v)
	}
	fmt.Fprint(out, ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
