// https://atcoder.jp/contests/abc197/tasks/abc197_e
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

func run(xcs [][]int) int {
	s := arrangexc(xcs)
	dp := []int{0, 0}
	pl, pr := 0, 0
	for i := 1; i < len(s); i++ {
		if s[i] == nil {
			continue
		}
		cl := s[i][0]
		cr := s[i][1]

		suml := dp[0]
		sumr := dp[1]
		dp[0] = min((cr-cl)+abs(pl-cr)+suml, (cr-cl)+abs(pr-cr)+sumr)
		dp[1] = min((cr-cl)+abs(pl-cl)+suml, (cr-cl)+abs(pr-cl)+sumr)

		pl, pr = cl, cr
	}

	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func arrangexc(xcs [][]int) [][]int {
	r := make([][]int, len(xcs)+2)
	var maxc int
	for _, v := range xcs {
		x := v[0]
		c := v[1]
		if r[c] == nil {
			r[c] = []int{x, x}
		} else {
			r[c][0] = min(r[c][0], x)
			r[c][1] = max(r[c][1], x)
		}
		maxc = max(maxc, c)
	}
	r[0] = []int{0, 0}
	r[maxc+1] = []int{0, 0}
	return r[:maxc+2]
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	xcs := make([][]int, n)
	for i := 0; i < n; i++ {
		xcs[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(xcs))
}
