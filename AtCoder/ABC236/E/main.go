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

	n := readInt()
	as := make([]int, n)
	avgArr := make([]int, n)
	for i := range as {
		as[i] = readInt()
		avgArr[i] = as[i] * 1000
	}

	low, high := 0, int(1e12+1)
	for low+1 < high {
		mid := low + (high-low)/2
		if favg(mid, avgArr) {
			low = mid
		} else {
			high = mid
		}
	}
	ansAvg := float64(low) / 1000

	low, high = 0, int(1e9+1)
	for low+1 < high {
		mid := low + (high-low)/2
		if fmed(mid, as) {
			low = mid
		} else {
			high = mid
		}
	}
	ansMed := low

	return fmt.Sprintf("%v %v", ansAvg, ansMed)
}

func favg(a int, as []int) bool {
	avgAs := make([]int, len(as))
	for i, v := range as {
		avgAs[i] = v - a
	}

	dp := make([][]int, len(as)+1)
	dp[0] = []int{0, 0}
	for i, a := range avgAs {
		n := dp[i][1]
		y := max(dp[i][0]+a, dp[i][1]+a)
		dp[i+1] = []int{n, y}
	}
	l := dp[len(as)]
	return (max(l[0], l[1]) > 0)
}

func fmed(a int, as []int) bool {
	medAs := make([]int, len(as))
	for i, v := range as {
		if v >= a {
			medAs[i] = 1
		}
	}

	cnt0, cnt1 := 0, 0
	flg := true
	for _, v := range medAs {
		if v > 0 {
			cnt1++
			flg = true
		} else if !flg {
			cnt0++
			flg = true
		} else {
			flg = false
		}
	}

	return cnt0 < cnt1
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
