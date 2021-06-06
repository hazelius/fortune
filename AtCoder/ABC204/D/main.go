// https://atcoder.jp/contests/ABC204/tasks/abc204_d
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

func run(n int, ts []int) int {
	sum := 0
	for _, v := range ts {
		sum += v
	}

	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		f(dp, ts[i])
	}
	// fmt.Println(dp)
	for i := sum/2 + sum%2; i < sum; i++ {
		if dp[i] {
			return i
		}
	}
	return sum
}

func f(dp []bool, t int) {
	res := make([]int, 0)
	for i, b := range dp {
		if b {
			res = append(res, i+t)
		}
	}
	for _, v := range res {
		dp[v] = true
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	ts := make([]int, n)
	for i := range ts {
		ts[i] = readInt()
	}
	fmt.Println(run(n, ts))
}
