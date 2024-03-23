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

type SortBy [][]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i][0] < a[j][0] }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	h, w, m := readInt(), readInt(), readInt()
	taxs := make([][]int, m)
	for i := range taxs {
		taxs[i] = []int{readInt(), readInt(), readInt()}
	}

	colmap := make(map[int]int)
	used1 := make(map[int]bool)
	used2 := make(map[int]bool)
	cnth, cntw := h, w

	for i := m - 1; i >= 0; i-- {
		t, a, x := taxs[i][0], taxs[i][1], taxs[i][2]
		if t == 1 {
			if used1[a] {
				continue
			}
			used1[a] = true
			colmap[x] += cntw
			cnth--
		} else {
			if used2[a] {
				continue
			}
			used2[a] = true
			colmap[x] += cnth
			cntw--
		}
	}
	colmap[0] += cntw * cnth
	for k, v := range colmap {
		if v == 0 {
			delete(colmap, k)
		}
	}

	ans := make([][]int, len(colmap))
	idx := 0
	for k, v := range colmap {
		ans[idx] = []int{k, v}
		idx++
	}

	sort.Sort(SortBy(ans))

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
