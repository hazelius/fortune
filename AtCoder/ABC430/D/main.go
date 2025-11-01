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
	xs := make([]int, n)
	for i := range xs {
		xs[i] = readInt()
	}
	tree := rbt.NewWithIntComparator()
	tree.Put(0, -1)
	ans := 0
	for _, x := range xs {
		fl, ok := tree.Floor(x)
		point := -1
		if ok {
			val := fl.Value.(int)
			key := fl.Key.(int)
			newVal := abs(key - x)
			if val < 0 || val > newVal {
				if val < 0 {
					ans += newVal
				} else {
					ans -= val - newVal
				}
				tree.Put(key, newVal)
			}
			point = newVal
		}
		cl, ok := tree.Ceiling(x)
		if ok {
			val := cl.Value.(int)
			key := cl.Key.(int)
			newVal := abs(key - x)
			if val < 0 || val > newVal {
				if val < 0 {
					ans += newVal
				} else {
					ans -= val - newVal
				}
				tree.Put(key, newVal)
			}
			if point < 0 || point > newVal {
				point = newVal
			}
		}
		ans += point
		tree.Put(x, point)
		fmt.Fprintln(out, ans)
	}

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
