// https://atcoder.jp/contests/abc217/tasks/abc217_d
// D - Cutting Woods

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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	l, q := readInt(), readInt()

	tree := rbt.NewWithIntComparator()
	tree.Put(0, 0)
	tree.Put(l, l)

	ans := make([]int, 0)
	for i := 0; i < q; i++ {
		c, x := readInt(), readInt()

		switch c {
		case 1:
			tree.Put(x, x)
		case 2:
			low, _ := tree.Floor(x)
			high, _ := tree.Ceiling(x)
			ans = append(ans, high.Value.(int)-low.Value.(int))
		}
	}

	ast := fmt.Sprintf("%v", ans)
	return ast[1 : len(ast)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
