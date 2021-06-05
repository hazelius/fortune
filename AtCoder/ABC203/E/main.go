// https://atcoder.jp/contests/abc203/tasks/abc203_e
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n, m int, xys [][]int) int {
	mapxy := make(map[int][]int)
	for _, xy := range xys {
		bp, ok := mapxy[xy[0]]
		if !ok {
			bp = make([]int, 0)
		}
		bp = append(bp, xy[1])
		mapxy[xy[0]] = bp
	}

	xs := make([]int, len(mapxy))
	i := 0
	for x := range mapxy {
		xs[i] = x
		i++
	}
	sort.Ints(xs)

	dp := make(map[int]bool)
	dp[n] = true
	for _, i := range xs {
		bp, ok := mapxy[i]
		if !ok {
			continue
		}
		// fmt.Printf("i:%v, bp:%v\n", i, bp)
		changes := make(map[int]bool)
		for _, y := range bp {
			if dp[y-1] || dp[y+1] {
				changes[y] = true
			} else if dp[y] {
				changes[y] = false
			}
		}
		for k, b := range changes {
			if b {
				dp[k] = b
			} else {
				delete(dp, k)
			}
		}
	}
	// fmt.Println(dp)

	return len(dp)
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := readInt(), readInt()
	xys := make([][]int, m)
	for i := range xys {
		xys[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(n, m, xys))
}
