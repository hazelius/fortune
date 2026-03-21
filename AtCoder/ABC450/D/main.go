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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	amap := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := readInt() % k
		amap[a] = true
	}

	if len(amap) == 1 {
		fmt.Fprint(out, 0)
		return
	}

	as := make([]int, len(amap))
	idx := 0
	for a := range amap {
		as[idx] = a
		idx++
	}

	sort.Ints(as)
	ans := 0
	for i, a := range as {
		tmp := (as[(i+1)%len(as)] + k - a) % k
		if tmp > ans {
			ans = tmp
		}
	}
	fmt.Fprint(out, k-ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
