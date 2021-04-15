// K - Stones
// https://atcoder.jp/contests/dp/tasks/dp_k

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

func run(n, k int, as []int) string {
	dp := make([]int, k+1)
	for i := 1; i <= k; i++ {
		for _, a := range as {
			if i-a < 0 {
				continue
			}
			if dp[i-a] == 0 {
				dp[i] = 1
			}
		}
	}
	if dp[k] == 1 {
		return "First"
	}
	return "Second"
}

func main() {
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, k, as))
}
