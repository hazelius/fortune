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
	uf := newUnionFind(n)
	hen := make(map[int]int)
	for i := 0; i < m; i++ {
		u, v := readInt()-1, readInt()-1
		if uf.same(u, v) {
			hen[uf.root(u)]++
			continue
		}

		uh := hen[uf.root(u)]
		uv := hen[uf.root(v)]
		uf.unite(u, v)
		rt := uf.root(u)
		hen[rt] = uh + uv + 1
	}

	for i := 0; i < n; i++ {
		rt := uf.root(i)
		if hen[rt] != uf.size(i) {
			fmt.Fprint(out, "No")
			return
		}
	}
	fmt.Fprint(out, "Yes")
}

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

func (u UnionFind) same(x, y int) bool {
	return u.root(x) == u.root(y)
}

func (u UnionFind) size(x int) int {
	return -u.node[u.root(x)]
}

func main() {
	run(os.Stdin, os.Stdout)
}
