// https://atcoder.jp/contests/abc214/tasks/abc214_d
// D - Sum of Maximum Weights

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

// Disjoint Set Union: Union Find Tree

type DSU struct {
	parentOrSize []int
	n            int
}

func newDsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	for i := range d.parentOrSize {
		d.parentOrSize[i] = -1
	}
	return &d
}

func (d DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

func (d DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

func (d DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d DSU) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

type Path struct {
	u int
	v int
	w int
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	pathes := make([]Path, n-1)
	for i := range pathes {
		pathes[i] = Path{u: readInt() - 1, v: readInt() - 1, w: readInt()}
	}

	sort.Slice(pathes, func(i, j int) bool { return pathes[i].w < pathes[j].w })

	d := newDsu(n)

	ans := 0
	for _, p := range pathes {
		us := d.Size(p.u)
		vs := d.Size(p.v)
		ans += us * vs * p.w
		d.Merge(p.u, p.v)
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
