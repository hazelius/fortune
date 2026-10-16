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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n, w := readInt(), readInt()

	dp := make([]map[int]int, 2)
	dp[0] = make(map[int]int)
	dp[0][0] = 0

	idx := 0
	for i := 0; i < n; i++ {
		a, b := readInt(), readInt()
		dp[1-idx] = make(map[int]int)
		for b2, a2 := range dp[idx] {
			dp[1-idx][b2] = max(dp[1-idx][b2], a2)
			v2, ok := dp[idx][b2]
			if b2+b <= w && ok {
				dp[1-idx][b2+b] = max(dp[1-idx][b2+b], v2+a)
			}
		}
		idx = 1 - idx
	}

	ans := 0
	for _, v := range dp[idx] {
		ans = max(ans, v)
	}
	fmt.Fprint(out, ans)
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
