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

	n, k := readInt(), readInt()
	pmap := make(map[int]int)
	for i := 0; i < n; i++ {
		p := readInt()
		pmap[p] = i
	}

	tree := rbt.NewWithIntComparator()
	ans := n
	for i := 1; i <= n; i++ {
		p := pmap[i]
		tree.Put(p, i)

		if i < k {
			continue
		}

		if i > k {
			tree.Remove(pmap[i-k])
		}

		l := tree.Left().Key.(int)
		r := tree.Right().Key.(int)

		tmp := r - l
		if tmp < ans {
			ans = tmp
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
