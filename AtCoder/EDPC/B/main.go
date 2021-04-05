package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n, k int, h []int) int {
	dp := make([]int, n+1)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j <= k; j++ {
			if j > i {
				break
			}
			dp[i] = min(dp[i], dp[i-j]+abs(h[i-j]-h[i]))
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
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

func main() {
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = readInt()
	}
	fmt.Println(run(n, k, h))
}
