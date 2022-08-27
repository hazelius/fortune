// E - Throwing the Die
// https://atcoder.jp/contests/abc266/tasks/abc266_e
// 期待値

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
	sc.Split(bufio.ScanWords)

	n := readInt()
	ans := 3.5
	for i := 1; i < n; i++ {
		ans = max(1, ans) + max(2, ans) + max(3, ans) + max(4, ans) + max(5, ans) + max(6, ans)
		ans /= 6.0
	}

	fmt.Fprintf(out, "%0.10f", ans)
}

func max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
