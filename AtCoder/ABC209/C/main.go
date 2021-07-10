package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner
var mod int = 1e9 + 7

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	cs := make([]int, n)
	for i := range cs {
		cs[i] = readInt()
	}

	sort.Ints(cs)
	ans := 1
	for i, c := range cs {
		if c-i <= 0 {
			return 0
		}
		ans = ans * (c - i) % mod
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
