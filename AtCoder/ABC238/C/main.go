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

	n := readInt()
	ans := 0
	for i := 1; i <= n; i *= 10 {
		t := i * 10
		if t > n {
			t = n + 1
		}
		t -= i
		ans += (t % mod) * ((t + 1) % mod) / 2
		ans %= mod
	}
	// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22    99 100 101 102 109 110
	// 1 2 3 4 5 6 7 8 9  1  2  3  4  5  6  7  8  9 10 11 12 13 ...90   1   2   3  10  11
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
