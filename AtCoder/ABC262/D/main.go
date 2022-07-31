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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	ans := 0
	for i := 1; i <= n; i++ {
		dp := make([][][]int, n+1)
		dp[0] = make([][]int, n+2)
		for j := range dp[0] {
			dp[0][j] = make([]int, i+1)
		}
		dp[0][0][0] = 1

		for j := 0; j < n; j++ {
			dp[j+1] = make([][]int, n+2)
			dp[j+1][0] = make([]int, i+1)
			for k := 0; k <= i; k++ {
				dp[j+1][k+1] = make([]int, i+1)
				for l := 0; l < i; l++ {
					dp[j+1][k][l] = (dp[j+1][k][l] + dp[j][k][l]) % mod
					if k != i {
						amari := (l + as[j]) % i
						dp[j+1][k+1][amari] = (dp[j+1][k+1][amari] + dp[j][k][l]) % mod
					}
				}
			}
		}
		ans = (ans + dp[n][i][0]) % mod
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
