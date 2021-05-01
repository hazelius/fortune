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

func run(n, d, h int, dhs [][]int) float64 {
	fh, fd := float64(h), float64(d)
	var ans float64
	for _, dh := range dhs {
		a := float64(dh[1]-h) / float64(dh[0]-d)
		b := fh - a*fd
		if ans < b {
			ans = b
		}
	}

	if ans < 0 {
		ans = 0
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	n, d, h := readInt(), readInt(), readInt()
	dhs := make([][]int, n)
	for i := range dhs {
		dhs[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(n, d, h, dhs))
}
