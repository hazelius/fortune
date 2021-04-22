// P - Independent Set
// https://atcoder.jp/contests/dp/tasks/dp_p

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var mod = 1000000007

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, xys [][]int) int {
	xymap := make(map[int][]int, n+1)
	for _, xy := range xys {
		x, y := xy[0], xy[1]
		xymap[x] = append(xymap[x], y)
		xymap[y] = append(xymap[y], x)
	}

	dp := make([][]int, n+1)
	f(0, 1, xymap, dp)
	// fmt.Println(dp)
	return (dp[1][0] + dp[1][1]) % mod
}

func f(p, c int, xymap map[int][]int, dp [][]int) {
	if dp[c] != nil {
		return
	}
	dp[c] = []int{1, 1}
	for _, j := range xymap[c] {
		if j == p {
			continue
		}
		f(c, j, xymap, dp)
		dp[c][0] = dp[c][0] * (dp[j][0] + dp[j][1]) % mod
		dp[c][1] = dp[c][1] * dp[j][0] % mod
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	xys := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		xys[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(n, xys))
}
