package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

var sc *bufio.Scanner

type wall struct {
	idx    int
	height int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ws := make([]wall, n)
	for i := range ws {
		ws[i] = wall{idx: i, height: readInt()}
	}

	sort.Slice(ws, func(i, j int) bool {
		return ws[i].height > ws[j].height
	})

	tree := rbt.NewWithIntComparator()
	ans := make([]int, n)

	for _, w := range ws {
		pre, ok := tree.Floor(w.idx)
		if !ok {
			ans[w.idx] = w.height * (w.idx + 1)
		} else {
			pidx := pre.Key.(int)
			v := ans[pidx]
			v += w.height * (w.idx - pidx)
			ans[w.idx] = v
		}
		tree.Put(w.idx, w.idx)
	}

	for _, v := range ans {
		fmt.Fprint(out, v+1, " ")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
