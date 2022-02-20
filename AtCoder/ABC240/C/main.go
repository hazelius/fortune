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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, x := readInt(), readInt()
	abs := make([][]int, n)
	for i := range abs {
		abs[i] = []int{readInt(), readInt()}
	}
	dp := make([]map[int]bool, n+1)
	dp[0] = make(map[int]bool)
	dp[0][0] = true
	for i, ab := range abs {
		dp[i+1] = make(map[int]bool)
		for k := range dp[i] {
			dp[i+1][k+ab[0]] = true
			dp[i+1][k+ab[1]] = true
		}
	}
	if dp[n][x] {
		return "Yes"
	}
	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
