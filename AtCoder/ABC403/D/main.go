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

	n, d := readInt(), readInt()
	if d == 0 {
		amap := make(map[int]int)
		for i := 0; i < n; i++ {
			a := readInt()
			amap[a]++
		}
		ans := 0
		for _, v := range amap {
			ans += v - 1
		}
		fmt.Fprint(out, ans)
		return
	}

	maxa := 0
	amap := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		if a > maxa {
			maxa = a
		}
		amap[a]++
	}

	ans := 0
	for i := 0; i < d; i++ {
		dp := [2][2]int{}
		idx := 0
		for j := i; j <= maxa; j += d {
			dp[1-idx][0] = amap[j] + min(dp[idx][0], dp[idx][1])
			dp[1-idx][1] = dp[idx][0]
			idx = 1 - idx
		}
		ans += min(dp[idx][0], dp[idx][1])
	}

	fmt.Fprint(out, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
