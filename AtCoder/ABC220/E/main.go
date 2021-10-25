// https://atcoder.jp/contests/abc220/tasks/abc220_e
// E - Distance on Large Perfect Binary Tree
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

	n, d := readInt(), readInt()

	bins := make([]int, n*2)
	for i := range bins {
		if i == 0 {
			bins[i] = 1
			continue
		}
		bins[i] = (bins[i-1] * 2) % mod
	}

	g := make([]int, n+1)
	for i := 1; i <= n; i++ {
		l := i - 1
		r := d - l
		if r >= 0 && r < i {
			li := max(l-1, 0)
			ri := max(r-1, 0)
			g[i] = (bins[li] * bins[ri]) % mod
			if l != r {
				g[i] = (g[i] * 2) % mod
			}
		}
		g[i] = (g[i] + g[i-1]) % mod
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = (f[i-1]*2 + g[i]) % mod
	}

	return (f[n] * 2) % mod
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
