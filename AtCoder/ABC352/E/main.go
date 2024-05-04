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

type kc struct {
	c  int
	as []int
}

type kcs []kc

func (a kcs) Len() int           { return len(a) }
func (a kcs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a kcs) Less(i, j int) bool { return a[i].c < a[j].c }

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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	vals := make(kcs, m)
	for i := 0; i < m; i++ {
		k, c := readInt(), readInt()
		as := make([]int, k)
		for i := range as {
			as[i] = readInt() - 1
		}
		vals[i] = kc{c, as}
	}

	sort.Sort(vals)

	ans := 0
	uf := newUnionFind(n)
	x := -1
	for _, v := range vals {
		i := 0
		for j := i + 1; j < len(v.as); j++ {
			if uf.same(v.as[i], v.as[j]) {
				continue
			}
			uf.unite(v.as[i], v.as[j])
			ans += v.c
			x = v.as[i]
			if uf.size(x) == n {
				fmt.Fprint(out, ans)
				return
			}
		}
	}

	if x == -1 || uf.size(x) < n {
		fmt.Fprint(out, -1)
		return

	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
