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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	bs := make([]int, n)
	for i := range bs {
		bs[i] = readInt()
	}
	cs := make([]int, n)
	for i := range cs {
		cs[i] = readInt()
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 4)
		for j := range dp[i] {
			dp[i][j] = -1 << 60
		}
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= 1; j++ {
			dp[i+1][1] = max(dp[i+1][1], dp[i][j]+as[i])
		}
		for j := 1; j <= 2; j++ {
			dp[i+1][2] = max(dp[i+1][2], dp[i][j]+bs[i])
		}
		for j := 2; j <= 3; j++ {
			dp[i+1][3] = max(dp[i+1][3], dp[i][j]+cs[i])
		}
	}

	fmt.Fprint(out, dp[n][3])

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
