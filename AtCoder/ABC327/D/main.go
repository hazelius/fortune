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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	as := make([]int, m)
	for i := range as {
		as[i] = readInt() - 1
	}
	bs := make([]int, m)
	for i := range bs {
		bs[i] = readInt() - 1
		if as[i] == bs[i] {
			fmt.Fprint(out, "No")
			return
		}
	}

	abmap := make(map[int]map[int]bool)
	for i, a := range as {
		b := bs[i]
		ab, ok := abmap[a]
		if !ok {
			ab = make(map[int]bool)
		}
		ab[b] = true
		abmap[a] = ab

		ba, ok := abmap[b]
		if !ok {
			ba = make(map[int]bool)
		}
		ba[a] = true
		abmap[b] = ba
	}

	xs := make([]int, n)
	for i := 0; i < n; i++ {
		if !f(i, xs, abmap) {
			fmt.Fprint(out, "No")
			return
		}
	}

	fmt.Fprint(out, "Yes")
}

func f(i int, xs []int, abmap map[int]map[int]bool) bool {
	ax := xs[i]
	if ax == 0 {
		ax = 1
		xs[i] = ax
	}

	ab, ok := abmap[i]
	if !ok {
		return true
	}
	for b := range ab {
		bx := xs[b]
		if bx == 0 {
			xs[b] = ax * -1
			if !f(b, xs, abmap) {
				return false
			}
		} else if bx == ax {
			return false
		}
	}

	return true
}

func main() {
	run(os.Stdin, os.Stdout)
}
