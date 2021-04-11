// https://atcoder.jp/contests/dp/tasks/dp_f

package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(s, t string) string {
	dp := make([][]int, len(s)+1)
	dp[0] = make([]int, len(t)+1)

	for i, sc := range s {
		dp[i+1] = make([]int, len(t)+1)
		for j, tc := range t {
			if sc == tc {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	i, j := len(s), len(t)
	ans := make([]byte, dp[i][j])
	for len := dp[i][j]; len > 0 && i > 0 && j > 0; {
		if s[i-1] == t[j-1] {
			ans[len-1] = s[i-1]
			len--
			i--
			j--
		} else if dp[i][j] == dp[i-1][j] {
			i--
		} else {
			j--
		}
	}

	return string(ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	s, t := readString(), readString()
	fmt.Println(run(s, t))
}
