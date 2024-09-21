package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type UnionFind struct {
	node []int
	vals [][]int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.node = make([]int, N)
	u.vals = make([][]int, N)
	for i := range u.node {
		u.node[i] = -1
		u.vals[i] = []int{i}
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

func (u UnionFind) rootKeys(x int) []int {
	r := u.root(x)
	return u.vals[r]
}

func (u UnionFind) unite(x, y int) {
	xr := u.root(x)
	yr := u.root(y)
	if xr == yr {
		return
	}
	u.node[yr] += u.node[xr]
	u.node[xr] = yr

	newKeys := make([]int, 0)
	xi, yi := 0, 0
	for i := 0; i < 10; i++ {
		xval, yval := -1, -1
		if xi < len(u.vals[xr]) {
			xval = u.vals[xr][xi]
		}
		if yi < len(u.vals[yr]) {
			yval = u.vals[yr][yi]
		}
		if xval < 0 && yval < 0 {
			break
		}

		if xval > yval {
			newKeys = append(newKeys, xval)
			xi++
		} else {
			newKeys = append(newKeys, yval)
			yi++
		}
	}
	u.vals[yr] = newKeys
}

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	uf := newUnionFind(n)

	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			u, v := readInt()-1, readInt()-1
			uf.unite(u, v)
		case 2:
			v, k := readInt()-1, readInt()-1
			keys := uf.rootKeys(v)
			if len(keys) <= k {
				fmt.Fprintln(out, -1)
				continue
			}
			fmt.Fprintln(out, keys[k]+1)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
