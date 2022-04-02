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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, k, x := readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	for i, a := range as {
		c := a / x
		if k < c {
			as[i] -= x * k
			k = 0
			break
		}
		k -= c
		as[i] = a % x
	}

	sort.Sort(sort.Reverse(sort.IntSlice(as)))
	for i := 0; i < k; i++ {
		if i >= len(as) || as[i] == 0 {
			break
		}
		as[i] = 0
	}

	ans := 0
	for _, a := range as {
		ans += a
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
