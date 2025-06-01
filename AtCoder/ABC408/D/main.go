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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		s := readString()

		dp := make([][]int, n)
		if s[0] == '0' {
			dp[0] = []int{0, 1, n + 1}
		} else {
			dp[0] = []int{1, 0, n + 1}
		}
		for j := 1; j < n; j++ {
			dp[j] = make([]int, 3)
			if s[j] == '0' {
				dp[j][0] = dp[j-1][0]
				dp[j][1] = min(dp[j-1][1], dp[j-1][0]) + 1
				dp[j][2] = min(dp[j-1][2], dp[j-1][1])
			} else {
				dp[j][0] = dp[j-1][0] + 1
				dp[j][1] = min(dp[j-1][1], dp[j-1][0])
				dp[j][2] = min(dp[j-1][2], dp[j-1][1]) + 1
			}
		}
		ans := min(min(dp[n-1][0], dp[n-1][1]), dp[n-1][2])
		fmt.Fprintln(out, ans)
	}
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
