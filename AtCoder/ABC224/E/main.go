// https://atcoder.jp/contests/abc224/tasks/abc224_e
// E - Integers on Grid
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

type Masu struct {
	i int
	r int
	c int
	a int
}

type MasuBy []*Masu

func (a MasuBy) Len() int           { return len(a) }
func (a MasuBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MasuBy) Less(i, j int) bool { return a[i].a > a[j].a }

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	h, w, n := readInt(), readInt(), readInt()
	masus := make([]*Masu, n)
	for i := range masus {
		masus[i] = &Masu{i: i, r: readInt(), c: readInt(), a: readInt()}
	}
	sort.Sort(MasuBy(masus))

	hmax := make([]int, h+1)
	wmax := make([]int, w+1)
	for i := range hmax {
		hmax[i] = -1
	}
	for i := range wmax {
		wmax[i] = -1
	}

	ans := make([]int, n)
	prea := -1
	halt := make([][]int, 0)
	walt := make([][]int, 0)
	for _, ms := range masus {
		if prea != ms.a {
			for _, v := range halt {
				hmax[v[0]] = max(v[1], hmax[v[0]])
			}
			for _, v := range walt {
				wmax[v[0]] = max(v[1], wmax[v[0]])
			}
			halt = make([][]int, 0)
			walt = make([][]int, 0)
		}
		prea = ms.a

		m := max(hmax[ms.r], wmax[ms.c]) + 1
		ans[ms.i] = m
		halt = append(halt, []int{ms.r, m})
		walt = append(walt, []int{ms.c, m})
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
