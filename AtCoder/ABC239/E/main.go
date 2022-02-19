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
var tops [][]int
var used []bool

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	xs := make([]int, n)
	for i := range xs {
		xs[i] = readInt()
	}

	abmap := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := readInt()-1, readInt()-1
		abmap[a] = append(abmap[a], b)
		abmap[b] = append(abmap[b], a)
	}

	tops = make([][]int, n)
	used = make([]bool, n)
	getTop(0, xs, abmap)

	ans := make([]int, q)
	for i := range ans {
		v, k := readInt()-1, readInt()-1
		ans[i] = tops[v][k]
	}

	ansstr := fmt.Sprintf("%d", ans)
	return ansstr[1 : len(ansstr)-1]
}

func getTop(i int, xs []int, abmap map[int][]int) []int {
	used[i] = true

	lst := []int{xs[i]}
	ab, ok := abmap[i]
	if !ok {
		tops[i] = lst
		return tops[i]
	}

	for _, b := range ab {
		if used[b] {
			continue
		}
		t := tops[b]
		if len(t) == 0 {
			t = getTop(b, xs, abmap)
		}
		lst = append(lst, t...)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lst)))
	if len(lst) > 20 {
		lst = lst[:20]
	}
	tops[i] = lst
	return tops[i]
}

func main() {
	fmt.Println(run(os.Stdin))
}
