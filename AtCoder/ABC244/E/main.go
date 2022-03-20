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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m, k, s, t, x := readInt(), readInt(), readInt(), readInt(), readInt(), readInt()
	uvmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		u, v := readInt(), readInt()
		uvmap[u] = append(uvmap[u], v)
		uvmap[v] = append(uvmap[v], u)
	}

	dp := make([][][]int, k+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	dp[0][s][0] = 1

	for i := 1; i <= k; i++ {
		for j, a := range dp[i-1] {
			if a[0] == 0 && a[1] == 0 {
				continue
			}

			for _, b := range uvmap[j] {
				flg := 0
				if b == x {
					flg = 1
				}
				dp[i][b][flg] = (dp[i][b][flg] + a[0]) % mod
				dp[i][b][1-flg] = (dp[i][b][1-flg] + a[1]) % mod
			}
		}
	}
	return dp[k][t][0]
}

func main() {
	fmt.Println(run(os.Stdin))
}
