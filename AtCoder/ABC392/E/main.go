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
	amari := make([][]int, 0)
	for i := 0; i < m; i++ {
		a, b := readInt()-1, readInt()-1
		if uf.same(a, b) {
			amari = append(amari, []int{i, a, b})
		} else {
			uf.unite(a, b)
		}
	}

	roots := make(map[int]bool)
	for i := 0; i < n; i++ {
		rt := uf.root(i)
		roots[rt] = true
	}

	if len(roots) < 2 {
		fmt.Fprint(out, 0)
		return
	}

	fmt.Fprintln(out, len(roots)-1)
	for _, am := range amari {
		rt := uf.root(am[1])
		for k := range roots {
			if k == rt {
				continue
			}
			if uf.same(rt, k) {
				continue
			}
			fmt.Fprintln(out, am[0]+1, am[2]+1, k+1)
			uf.unite(rt, k)
			delete(roots, k)
			break
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
