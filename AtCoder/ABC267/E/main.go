package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
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

	n, m := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	mapn := make(map[int][]int)
	for i := 0; i < m; i++ {
		u, v := readInt()-1, readInt()-1
		mapn[u] = append(mapn[u], v)
		mapn[v] = append(mapn[v], u)
	}

	grvs := make([]int, n)
	tree := rbt.NewWithIntComparator()
	for a, vs := range mapn {
		grv := 0
		for _, v := range vs {
			grv += as[v]
		}
		grvs[a] = grv
		appendNode(a, grv, tree)
	}

	ans := 0
	removed := make([]bool, n)

	for {
		v := tree.Left()
		if v == nil {
			break
		}
		idxs := v.Value.([]int)
		idx := getmax(idxs, grvs)

		removeNode(idx, v.Key.(int), tree)
		removed[idx] = true

		if ans < grvs[idx] {
			ans = grvs[idx]
		}

		a := as[idx]
		for _, j := range mapn[idx] {
			if removed[j] {
				continue
			}

			g := grvs[j]
			removeNode(j, g, tree)

			g -= a
			grvs[j] = g
			appendNode(j, g, tree)
		}
	}

	fmt.Fprint(out, ans)
}

func getmax(idxs, as []int) int {
	mxa := -1
	ret := 0
	for _, idx := range idxs {
		a := as[idx]
		if mxa < a {
			mxa = a
			ret = idx
		}
	}
	return ret
}

func appendNode(v, k int, tree *rbt.Tree) {
	n, ok := tree.Get(k)
	idxs := []int{v}
	if ok {
		idxs = append(n.([]int), v)
	}
	tree.Put(k, idxs)
}

func removeNode(v, k int, tree *rbt.Tree) {
	ns, _ := tree.Get(k)
	nsa := ns.([]int)
	if len(nsa) == 1 {
		tree.Remove(k)
		return
	}

	newn := make([]int, len(nsa)-1)
	plsidx := 0
	for _, ps := range nsa {
		if ps != v {
			newn[plsidx] = ps
			plsidx++
		}
	}
	tree.Put(k, newn)
}

func main() {
	run(os.Stdin, os.Stdout)
}
