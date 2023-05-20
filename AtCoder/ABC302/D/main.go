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

	n, m, d := readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	tree := rbt.NewWithIntComparator()

	for i := 0; i < m; i++ {
		b := readInt()
		tree.Put(b, b)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(as)))

	for _, a := range as {
		b, ok := tree.Floor(a + d)
		if !ok {
			continue
		}
		bval := b.Value.(int)
		if bval-a < -d {
			continue
		}

		fmt.Fprint(out, a+bval)
		return
	}

	fmt.Fprint(out, -1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
