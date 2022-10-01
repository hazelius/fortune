// https://atcoder.jp/contests/abc253/tasks/abc253_c

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

	q := readInt()

	tree := rbt.NewWithIntComparator()
	xmap := make(map[int]int)
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x := readInt()
			tree.Put(x, true)
			xmap[x]++
		case 2:
			x, c := readInt(), readInt()
			cnt, ok := xmap[x]
			if ok {
				cnt -= c
				if cnt <= 0 {
					cnt = 0
					tree.Remove(x)
				}
				xmap[x] = cnt
			}
		case 3:
			right := tree.Right()
			left := tree.Left()
			fmt.Fprintln(out, right.Key.(int)-left.Key.(int))
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
