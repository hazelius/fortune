package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ss := make([]int, n)
	for i := range ss {
		ss[i] = readInt()
	}
	ts := make([]int, n)
	for i := range ts {
		ts[i] = readInt()
	}

	dp := make([]int, n)
	for i := 0; i < 2*n; i++ {
		ci := i % n
		if dp[ci] == 0 {
			dp[ci] = ts[ci]
		} else {
			dp[ci] = min(ts[ci], dp[ci])
		}

		ni := (i + 1) % n
		nv := dp[ci] + ss[ci]
		if dp[ni] == 0 {
			dp[ni] = nv
		} else {
			dp[ni] = min(nv, dp[ni])
		}
	}

	ans := strings.Replace(fmt.Sprintf("%v", dp), " ", "\n", -1)
	return ans[1 : len(ans)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
