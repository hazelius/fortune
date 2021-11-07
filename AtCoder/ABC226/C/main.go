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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ts := make([]int, n)
	ks := make([][]int, n)
	for i := range ts {
		ts[i] = readInt()
		k := readInt()
		ks[i] = make([]int, k)
		for j := 0; j < k; j++ {
			ks[i][j] = readInt() - 1
		}
	}

	ans := 0
	needs := make(map[int]bool)
	needs[n-1] = true
	for i := n - 1; i >= 0; i-- {
		if !needs[i] {
			continue
		}
		ans += ts[i]
		for _, v := range ks[i] {
			needs[v] = true
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
