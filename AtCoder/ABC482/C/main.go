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

	_, _, n, q := readInt(), readInt(), readInt(), readInt()
	xs, ys := make([]int, n), make([]int, n)
	for i := range xs {
		xs[i], ys[i] = readInt()-1, readInt()-1
	}
	sort.Ints(xs)
	sort.Ints(ys)

	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			u, d := readInt()-1, readInt()-1
			idxu := sort.SearchInts(xs, u)
			idxd := sort.SearchInts(xs, d+1)
			fmt.Fprintln(out, idxd-idxu)
		case 2:
			l, r := readInt()-1, readInt()-1
			idxl := sort.SearchInts(ys, l)
			idxr := sort.SearchInts(ys, r+1)

			fmt.Fprintln(out, idxr-idxl)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
