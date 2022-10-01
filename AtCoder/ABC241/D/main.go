// https://atcoder.jp/contests/abc241/tasks/abc241_d
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
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x := readInt()
			v, ok := tree.Get(x)
			if ok {
				tree.Put(x, v.(int)+1)
			} else {
				tree.Put(x, 1)
			}
		case 2:
			x, k := readInt(), readInt()
			ans := -1
			for i := 0; i < k; i++ {
				a, ok := tree.Floor(x)
				if !ok {
					ans = -1
					break
				}
				cnt := a.Value.(int)
				i += cnt - 1
				x = a.Key.(int)
				ans = x
				x--
			}
			fmt.Fprintln(out, ans)
		case 3:
			x, k := readInt(), readInt()
			ans := -1
			for i := 0; i < k; i++ {
				a, ok := tree.Ceiling(x)
				if !ok {
					ans = -1
					break
				}
				cnt := a.Value.(int)
				i += cnt - 1
				x = a.Key.(int)
				ans = x
				x++
			}
			fmt.Fprintln(out, ans)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
