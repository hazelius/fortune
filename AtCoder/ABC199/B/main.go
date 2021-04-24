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

func run(n int, as, bs []int) int {
	maxa := 0
	for _, v := range as {
		maxa = max(maxa, v)
	}
	minb := bs[0]
	for _, v := range bs {
		minb = min(minb, v)
	}
	ans := 0
	if minb >= maxa {
		ans = minb - maxa + 1
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	bs := make([]int, n)
	for i := range bs {
		bs[i] = readInt()
	}
	fmt.Println(run(n, as, bs))
}
