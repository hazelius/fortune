// Q - Flowers
// https://atcoder.jp/contests/dp/tasks/dp_q
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

func run(n int, hs, as []int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		h := hs[i]
		a := as[i]
		maxa := 0
		for _, dpha := range dp {
			if dpha == nil || h < dpha[0] {
				continue
			}
			if maxa < dpha[1] {
				maxa = dpha[1]
			}
		}
		dp[i] = []int{h, maxa + a}
	}

	ans := 0
	for _, dpha := range dp {
		if dpha == nil {
			continue
		}
		if ans < dpha[1] {
			ans = dpha[1]
		}
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	hs := make([]int, n)
	for i := range hs {
		hs[i] = readInt()
	}
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, hs, as))
}
