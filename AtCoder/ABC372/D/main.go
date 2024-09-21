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

	n := readInt()
	hs := make([]int, n)

	for i := range hs {
		h := readInt()
		hs[i] = h
	}

	tree := rbt.NewWithIntComparator()

	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		ans[i] = tree.Size()
		h := hs[i]
		for {
			cl, ok := tree.Floor(h)
			if !ok {
				break
			}
			tree.Remove(cl.Key)
		}
		tree.Put(h, h)
	}

	str := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, str[1:len(str)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
