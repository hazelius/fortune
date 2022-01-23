package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var mx int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt() * 2
	as := make([][]int, n)
	for i := range as {
		as[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			as[i][j] = readInt()
		}
	}

	mx = 0
	used := make([]bool, n)
	f(0, 0, used, as)

	return mx
}

func f(a, sum int, used []bool, as [][]int) {
	if a == len(as) {
		if mx < sum {
			mx = sum
		}
		return
	}
	if used[a] {
		f(a+1, sum, used, as)
		return
	}
	used[a] = true

	for i := a + 1; i < len(as); i++ {
		if !used[i] {
			used[i] = true
			f(a+1, sum^as[a][i], used, as)
			used[i] = false
		}
	}
	used[a] = false
}

func main() {
	fmt.Println(run(os.Stdin))
}
