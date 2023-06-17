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
	dp := make([][]int, n+1)
	dp[0] = make([]int, 2)
	for i := 0; i < n; i++ {
		x, y := readInt(), readInt()
		dp[i+1] = dp[i]
		if x == 1 {
			dp[i+1][1] = max(dp[i][1], dp[i][0]+y)
		} else {
			v := max(dp[i][0], dp[i][0]+y)
			dp[i+1][0] = max(v, dp[i][1]+y)
		}
	}

	fmt.Fprint(out, max(dp[n][0], dp[n][1]))
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
