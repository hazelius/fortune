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

func run(n, m int, xys [][]int) int {
	mapxy := make(map[int][]bool)
	for _, xy := range xys {
		bp, ok := mapxy[xy[0]]
		if !ok {
			bp = make([]bool, 2*n+1)
		}
		bp[xy[1]] = true
		mapxy[xy[0]] = bp
	}

	dp := make([]bool, 2*n+1)
	dp[n] = true
	for i := 1; i < 2*n+1; i++ {
		bp, ok := mapxy[i]
		if !ok {
			continue
		}
		// fmt.Printf("i:%v, bp:%v\n", i, bp)
		newdp := make([]bool, 2*n+1)
		for j, b := range bp {
			if !b {
				newdp[j] = dp[j]
				continue
			}
			if j > 0 && dp[j-1] {
				newdp[j] = true
				continue
			}
			if j < 2*n && dp[j+1] {
				newdp[j] = true
			}
		}
		dp = newdp
	}
	fmt.Println(dp)

	ans := 0
	for _, v := range dp {
		if v {
			ans++
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, m := readInt(), readInt()
	xys := make([][]int, m)
	for i := range xys {
		xys[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(n, m, xys))
}
