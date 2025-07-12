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

	n, m := readInt(), readInt()
	xmap := make(map[int]bool)
	for i := 0; i < n; i++ {
		x := readInt()
		xmap[x] = true
	}
	if len(xmap) <= m {
		fmt.Fprint(out, 0)
		return
	}

	xs := make([]int, len(xmap))
	idx := 0
	for k := range xmap {
		xs[idx] = k
		idx++
	}
	sort.Ints(xs)

	ls := make([]int, len(xs)-1)
	for i := range ls {
		ls[i] = xs[i+1] - xs[i]
	}
	sort.Ints(ls)

	ans := 0
	for i := 0; i < len(xs)-m; i++ {
		ans += ls[i]
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
