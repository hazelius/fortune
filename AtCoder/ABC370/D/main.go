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

type wall struct {
	x int
	y int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	h, w, q := readInt(), readInt(), readInt()

	wmap := make(map[wall]bool)

	htrees := make([]*rbt.Tree, w)
	for i := range htrees {
		htrees[i] = rbt.NewWithIntComparator()
		for j := 0; j < h; j++ {
			appendNode(j, htrees[i])
		}
	}

	wtrees := make([]*rbt.Tree, h)
	for i := range wtrees {
		wtrees[i] = rbt.NewWithIntComparator()
		for j := 0; j < w; j++ {
			appendNode(j, wtrees[i])
		}
	}

	for i := 0; i < q; i++ {
		w := wall{readInt() - 1, readInt() - 1}
		if !wmap[w] {
			wmap[w] = true
			removeNode(w.x, htrees[w.y])
			removeNode(w.y, wtrees[w.x])
			continue
		}
		val, ok := htrees[w.y].Ceiling(w.x)
		if ok {
			bomb(val.Key.(int), w.y, htrees, wtrees, wmap)
		}
		val, ok = htrees[w.y].Floor(w.x)
		if ok {
			bomb(val.Key.(int), w.y, htrees, wtrees, wmap)
		}
		val, ok = wtrees[w.x].Ceiling(w.y)
		if ok {
			bomb(w.x, val.Key.(int), htrees, wtrees, wmap)
		}
		val, ok = wtrees[w.x].Floor(w.y)
		if ok {
			bomb(w.x, val.Key.(int), htrees, wtrees, wmap)
		}
	}

	fmt.Fprint(out, w*h-len(wmap))
}

func bomb(x, y int, htrees, wtrees []*rbt.Tree, wmap map[wall]bool) {
	wmap[wall{x, y}] = true
	removeNode(x, htrees[y])
	removeNode(y, wtrees[x])
}

func appendNode(k int, tree *rbt.Tree) {
	n, ok := tree.Get(k)
	if !ok {
		tree.Put(k, 1)
		return
	}

	v := n.(int) + 1
	tree.Put(k, v)
}

func removeNode(k int, tree *rbt.Tree) {
	ns, ok := tree.Get(k)
	if !ok {
		return
	}
	v := ns.(int)
	if v == 1 {
		tree.Remove(k)
		return
	}
	v--

	tree.Put(k, v)
}

func main() {
	run(os.Stdin, os.Stdout)
}
