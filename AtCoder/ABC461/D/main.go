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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		s := readString()

		xs := make([]int, n)
		for i := range xs {
			xs[i] = readInt()
		}
		ys := make([]int, n-1)
		for i := range ys {
			ys[i] = readInt()
		}

		dp := make([][]int, n)
		for idx, c := range s {
			dp[idx] = make([]int, 2)
			if idx > 0 {
				dp[idx][0] = max(dp[idx-1][0], dp[idx-1][1]+ys[idx-1])
				dp[idx][1] = max(dp[idx-1][0], dp[idx-1][1])
			}
			if c == 'S' {
				dp[idx][1] -= xs[idx]
			} else {
				dp[idx][0] -= xs[idx]
			}
		}

		fmt.Fprintln(out, max(dp[n-1][0], dp[n-1][1]))
	}

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
