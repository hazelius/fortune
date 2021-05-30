package main

import (
	"bufio"
	"fmt"
	"math"
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

func run(n, k int, sq [][]int) int {
	rg := k * k
	center := (rg / 2) + (rg % 2) - 1
	ans := math.MaxInt64
	// fmt.Println(center)
	for i := 0; i <= n-k; i++ {
		for j := 0; j <= n-k; j++ {
			ary := make([]int, rg)
			for i2 := range ary {
				ary[i2] = sq[i+i2/k][j+i2%k]
			}
			// fmt.Println(ary)
			sort.Ints(ary)
			if ary[center] < ans {
				ans = ary[center]
			}
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	sq := make([][]int, n)
	for i := range sq {
		sq[i] = make([]int, n)
		for j := range sq[i] {
			sq[i][j] = readInt()
		}
	}
	fmt.Println(run(n, k, sq))
}
