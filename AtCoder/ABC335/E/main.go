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

type node struct {
	idx int
	num int
}

type Nodes []node

func (a Nodes) Len() int           { return len(a) }
func (a Nodes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nodes) Less(i, j int) bool { return a[i].num < a[j].num }

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	as := make(Nodes, n)
	for i := range as {
		as[i] = node{i, readInt()}
	}

	uf := newUnionFind(n)
	uvmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		u, v := readInt()-1, readInt()-1
		if as[u].num == as[v].num {
			uf.unite(u, v)
			continue
		}
		if as[u].num > as[v].num {
			u, v = v, u
		}

		ms, ok := uvmap[u]
		if !ok {
			ms = []int{v}
		} else {
			ms = append(ms, v)
		}
		uvmap[u] = ms
	}

	dp := make([]int, n)
	dp[uf.root(0)] = 1

	sort.Sort(as)

	for _, a := range as {
		uvs, ok := uvmap[a.idx]
		if !ok {
			continue
		}
		u := uf.root(a.idx)
		if dp[u] <= 0 {
			continue
		}
		for _, v := range uvs {
			v = uf.root(v)
			point := dp[u] + 1
			if point > dp[v] {
				dp[v] = point
			}
		}
	}

	fmt.Fprint(out, dp[uf.root(n-1)])
}

func main() {
	run(os.Stdin, os.Stdout)
}
