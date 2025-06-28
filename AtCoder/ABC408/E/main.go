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
	uvws := make([][]int, m)
	for i := range uvws {
		uvws[i] = []int{readInt() - 1, readInt() - 1, readInt()}
	}

	ans := 0
	for i := 29; i >= 0; i-- {
		uf := newUnionFind(n)
		for _, uvw := range uvws {
			w := uvw[2]
			if (w>>i | ans>>i) == (ans >> i) {
				uf.unite(uvw[0], uvw[1])
			}
		}
		if !uf.same(0, n-1) {
			ans |= 1 << i
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
