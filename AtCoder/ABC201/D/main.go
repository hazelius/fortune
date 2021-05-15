package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(h, w int, as []string) string {
	dp := make([][][]int, h)
	for i := h - 1; i >= 0; i-- {
		dp[i] = make([][]int, w)
		for j := w - 1; j >= 0; j-- {
			t := (i + j) % 2
			if i == h-1 && j == w-1 {
				dp[i][j] = []int{0, 0}
			} else if i == h-1 {
				dp[i][j] = []int{dp[i][j+1][0], dp[i][j+1][1]}
			} else if j == w-1 {
				dp[i][j] = []int{dp[i+1][j][0], dp[i+1][j][1]}
			} else {
				if dp[i+1][j][1-t]-dp[i+1][j][t] < dp[i][j+1][1-t]-dp[i][j+1][t] {
					dp[i][j] = []int{dp[i][j+1][0], dp[i][j+1][1]}
				} else {
					dp[i][j] = []int{dp[i+1][j][0], dp[i+1][j][1]}
				}
			}

			if i == 0 && j == 0 {
				continue
			}

			if as[i][j] == '+' {
				dp[i][j][t]++
			} else {
				dp[i][j][t]--
			}
		}
	}

	if dp[0][0][0] < dp[0][0][1] {
		return "Takahashi"
	} else if dp[0][0][0] > dp[0][0][1] {
		return "Aoki"
	}
	return "Draw"
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	h, w := readInt(), readInt()
	as := make([]string, h)
	for i := range as {
		as[i] = readString()
	}
	fmt.Println(run(h, w, as))
}
