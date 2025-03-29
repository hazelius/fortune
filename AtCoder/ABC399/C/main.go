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

func (u UnionFind) same(x, y int) bool {
	return u.root(x) == u.root(y)
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
	ans := 0
	for i := 0; i < m; i++ {
		a, b := readInt()-1, readInt()-1
		if uf.same(a, b) {
			ans++
		} else {
			uf.unite(a, b)
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
