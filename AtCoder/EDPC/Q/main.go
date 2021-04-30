// Q - Flowers
// https://atcoder.jp/contests/dp/tasks/dp_q
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SegmentTree struct {
	data []int
	n    int
}

func NewSegmentTree(n int) *SegmentTree {
	segtree := new(SegmentTree)
	segtree.n = 1
	for segtree.n < n {
		segtree.n *= 2
	}
	segtree.data = make([]int, segtree.n*2-1)
	return segtree
}

func (s *SegmentTree) query(begin, end, idx, a, b int) int {
	if b <= begin || end <= a {
		// クエリと関係のない区間
		return 0
	}
	if begin <= a && b <= end {
		// 全体がクエリの対象になる区間
		return s.data[idx]
	}
	// 一部がクエリの対象にならない場合は子に尋ねる
	v1 := s.query(begin, end, idx*2+1, a, (a+b)/2)
	v2 := s.query(begin, end, idx*2+2, (a+b)/2, b)
	if v1 > v2 {
		return v1
	}
	return v2
}

func (segtree *SegmentTree) Query(begin, end int) int {
	return segtree.query(begin, end, 0, 0, segtree.n)
}

func (s *SegmentTree) Update(idx, x int) {
	idx += s.n - 1
	s.data[idx] = x
	for 0 < idx {
		idx = (idx - 1) / 2
		lval := s.data[idx*2+1]
		rval := s.data[idx*2+2]
		if lval > rval {
			s.data[idx] = lval
		} else {
			s.data[idx] = rval
		}
	}
}

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, hs, as []int) int {
	dp := NewSegmentTree(n)
	for i := 0; i < n; i++ {
		h := hs[i] - 1
		a := as[i]

		maxa := dp.Query(0, h)
		dp.Update(h, maxa+a)
	}

	return dp.Query(0, n)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	hs := make([]int, n)
	for i := range hs {
		hs[i] = readInt()
	}
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, hs, as))
}
