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

func run(n int, powers [][]int) int {
	ans, l, r := 0, 0, 1000000001
	for l+1 < r {
		ans = l + (r-l)/2
		for _, v := range powers {

		}
	}

	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	powers := make([][]int, n)
	for i := range powers {
		powers[i] = []int{readInt(), readInt(), readInt(), readInt(), readInt()}
	}
	fmt.Println(run(n, powers))
}
