// https://atcoder.jp/contests/dp/tasks/dp_h
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(h, w int, ahw []string) int {
	dp := make([][]int, h)
	for i, row := range ahw {
		dp[i] = make([]int, w)
		if i == 0 {
			dp[i][0] = 1
		}
		for j, cell := range row {
			if cell == '#' {
				continue
			}

			if j > 0 {
				if row[j-1] == '.' {
					dp[i][j] += dp[i][j-1]
				}
			}

			if i > 0 {
				if ahw[i-1][j] == '.' {
					dp[i][j] = (dp[i][j] + dp[i-1][j]) % (1000000007)
				}
			}
		}
	}

	return dp[h-1][w-1]
}

func f(x int, dp map[int]int, links map[int][]int) int {
	if v, ok := dp[x]; ok {
		return v
	}
	dp[x] = 0

	l, _ := links[x]
	for _, y := range l {
		dp[x] = max(dp[x], f(y, dp, links)+1)
	}
	return dp[x]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	h, w := readInt(), readInt()
	ahw := make([]string, h)
	for i := range ahw {
		ahw[i] = readString()
	}
	fmt.Println(run(h, w, ahw))
}
