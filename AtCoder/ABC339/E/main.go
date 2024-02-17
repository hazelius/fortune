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

	n, d := readInt(), readInt()
	as := make([]int, n)
	mx := 0
	for i := range as {
		as[i] = readInt()
		mx = max(mx, as[i])
	}

	dp := NewSegmentTree(mx+1, 0, func(a, b int) int { return max(a, b) })

	for _, a := range as {
		dp.Update(a, dp.Query(max(0, a-d), min(mx, a+d)+1)+1)
	}

	fmt.Fprint(out, dp.nodes[1])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}

type SegmentTree struct {
	size  int
	nodes []int
	f     func(x1, x2 int) int
	inf   int
}

func NewSegmentTree(n, inf int, f func(x1, x2 int) int) (st *SegmentTree) {
	st = new(SegmentTree)
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.nodes = make([]int, 2*st.size)
	for i := range st.nodes {
		st.nodes[i] = inf
	}
	st.inf = inf
	st.f = f
	return st
}

func (seg *SegmentTree) QueryRecursively(a, b, k, l, r int) int {
	// [a, b)と[l, r)が交差しない
	if a >= r || b <= l {
		return seg.inf
	}

	// [a, b)が[l, r)を完全に含んでいる
	if a <= l && b >= r {
		return seg.nodes[k]
	}

	vl := seg.QueryRecursively(a, b, 2*k, l, (l+r)/2)
	vr := seg.QueryRecursively(a, b, 2*k+1, (l+r)/2, r)
	return seg.f(vl, vr)
}

func (seg *SegmentTree) Query(l, r int) int {
	return seg.QueryRecursively(l, r, 1, 0, seg.size)
}

func (seg *SegmentTree) Update(k, x int) {
	k += seg.size
	seg.nodes[k] = seg.f(seg.nodes[k], x)
	for k > 1 {
		k /= 2
		seg.nodes[k] = seg.f(seg.nodes[k*2], seg.nodes[k*2+1])
	}
}
