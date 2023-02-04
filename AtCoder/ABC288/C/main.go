package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type point struct {
	a int
	b int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	abmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		abmap[a] = append(abmap[a], b)
		abmap[b] = append(abmap[b], a)
	}

	ans := 0
	used := make(map[int]bool)
	usedP := make(map[point]bool)
	for i := 1; i <= n; i++ {
		p, h := f(i, abmap, used, usedP)
		if p > 0 {
			sa := h - (p - 1)
			if sa > 0 {
				ans += sa
			}
		}
	}

	fmt.Fprint(out, ans)
}

func f(a int, abmap map[int][]int, used map[int]bool, usedP map[point]bool) (int, int) {
	if used[a] {
		return 0, 0
	}
	used[a] = true

	p, h := 1, 0
	ab, ok := abmap[a]
	if !ok {
		return p, h
	}

	for _, b := range ab {
		if usedP[point{a, b}] {
			continue
		}
		usedP[point{a, b}] = true
		usedP[point{b, a}] = true

		h++
		pt, ht := f(b, abmap, used, usedP)
		p += pt
		h += ht
	}

	return p, h
}

func main() {
	run(os.Stdin, os.Stdout)
}
