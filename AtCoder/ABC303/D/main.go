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

	x, y, z := readInt(), readInt(), readInt()
	s := readString()

	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	dp[1][0] = int(1e9)

	for i, c := range s {
		if c == 'a' {
			dp[0][i+1] = min(dp[0][i]+x, dp[1][i]+z+x)
			dp[1][i+1] = min(dp[1][i]+y, dp[0][i]+z+y)
		} else {
			dp[0][i+1] = min(dp[0][i]+y, dp[1][i]+z+y)
			dp[1][i+1] = min(dp[1][i]+x, dp[0][i]+z+x)
		}
	}

	fmt.Fprint(out, min(dp[0][len(s)], dp[1][len(s)]))
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
