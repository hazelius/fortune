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

	n, k := readInt(), readInt()
	ps := make([]int, n)
	for i := range ps {
		ps[i] = readInt()
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	tree := rbt.NewWithIntComparator()

	for i, p := range ps {
		mnt, ok := tree.Ceiling(p)
		val := make(map[int]bool)
		if ok {
			val = mnt.Value.(map[int]bool)
			tree.Remove(mnt.Key)
		}

		val[p-1] = true
		if len(val) < k {
			tree.Put(p, val)
		} else {
			for k := range val {
				ans[k] = i + 1
			}
		}
	}
	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
