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

	n, h, m := readInt(), readInt(), readInt()
	abs := make([][]int, n)
	for i := range abs {
		abs[i] = []int{readInt(), readInt()}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 3001)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][h] = m

	ans := 0
	for i, ab := range abs {
		for hp, mp := range dp[i] {
			tmp := mp - ab[1]
			if hp+ab[0] < len(dp[i]) {
				tmp = max(tmp, dp[i][hp+ab[0]])
			}
			dp[i+1][hp] = max(dp[i+1][hp], tmp)
			if dp[i+1][hp] >= 0 {
				ans = i + 1
			}
		}
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
