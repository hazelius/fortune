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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	tree := rbt.NewWithIntComparator()
	for i := 0; i < n; i++ {
		a := readInt()
		appendNode(a, tree)
	}

	bs := make([]int, m)
	for i := range bs {
		bs[i] = readInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(bs)))

	ans := 0
	for _, b := range bs {
		a, ok := tree.Ceiling(b)
		if !ok {
			fmt.Fprint(out, -1)
			return
		}
		ans += a.Key.(int)
		removeNode(a.Key.(int), tree)
	}

	fmt.Fprint(out, ans)
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
		fmt.Println(k)
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
