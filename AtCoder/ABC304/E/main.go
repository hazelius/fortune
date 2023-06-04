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

type pos struct {
	a int
	b int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()

	uf := newUnionFind(n)
	for i := 0; i < m; i++ {
		u, v := readInt()-1, readInt()-1
		uf.unite(u, v)
	}

	k := readInt()
	pmap := make(map[pos]bool)
	for i := 0; i < k; i++ {
		x, y := readInt()-1, readInt()-1
		xr := uf.root(x)
		yr := uf.root(y)
		if xr > yr {
			xr, yr = yr, xr
		}
		pmap[pos{xr, yr}] = true
	}

	qcnt := readInt()
	for i := 0; i < qcnt; i++ {
		p, q := readInt()-1, readInt()-1
		pr := uf.root(p)
		qr := uf.root(q)
		if pr > qr {
			pr, qr = qr, pr
		}
		if pmap[pos{pr, qr}] {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
