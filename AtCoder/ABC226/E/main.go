package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var mod = 998244353

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

func (u UnionFind) size(x int) int {
	return -u.node[u.root(x)]
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	mapuv := make(map[int][]int)
	uf := newUnionFind(n)

	for i := 0; i < m; i++ {
		u, v := readInt()-1, readInt()-1
		mapuv[u] = append(mapuv[u], v)
		mapuv[v] = append(mapuv[v], u)
		uf.unite(u, v)
	}

	unions := make(map[int]int)
	for i := 0; i < n; i++ {
		root := uf.root(i)
		uv, ok := mapuv[i]
		if ok {
			unions[root] += len(uv)
		} else {
			unions[root] += 0
		}
	}

	for k, v := range unions {
		unions[k] = v / 2
	}

	ans := 1
	for r, edges := range unions {
		size := uf.size(r)
		if size == 0 {
			continue
		}
		if size != edges {
			ans = 0
		} else {
			ans = (ans * 2) % mod
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
