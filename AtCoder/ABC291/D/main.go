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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	abs := make([][]int, n)
	for i := range abs {
		abs[i] = []int{readInt(), readInt()}
	}

	dp := make([][]int, n)
	dp[0] = []int{1, 1}
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 2)
		for preIdx, prev := range abs[i-1] {
			for curIdx, v := range abs[i] {
				if prev != v {
					dp[i][curIdx] = (dp[i][curIdx] + dp[i-1][preIdx]) % mod
				}
			}
		}
	}

	fmt.Fprint(out, (dp[n-1][0]+dp[n-1][1])%mod)
}

func main() {
	run(os.Stdin, os.Stdout)
}
