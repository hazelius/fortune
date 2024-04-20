package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type UnionFind struct {
	node []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.node = make([]int, N)
	for i := range u.node {
		u.node[i] = -1
	}
	return u
}

func (u UnionFind) root(x int) int {
	if u.node[x] < 0 {
		return x
	}
	u.node[x] = u.root(u.node[x])
	return u.node[x]
}

func (u UnionFind) unite(x, y int) {
	xr := u.root(x)
	yr := u.root(y)
	if xr == yr {
		return
	}
	u.node[yr] += u.node[xr]
	u.node[xr] = yr
}

func (u UnionFind) size(x int) int {
	return -u.node[u.root(x)]
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
	uf := newUnionFind(n)

	for i := 0; i < m; i++ {
		a, b := readInt()-1, readInt()-1
		uf.unite(a, b)
	}

	ans := 0
	used := make(map[int]bool)
	for i := 0; i < n; i++ {
		root := uf.root(i)
		if used[root] {
			continue
		}
		used[root] = true

		size := uf.size(i)
		ans += size * (size - 1) / 2
	}

	fmt.Fprint(out, ans-m)
}

func main() {
	run(os.Stdin, os.Stdout)
}
