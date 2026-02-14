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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n, d := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	tree := rbt.NewWithIntComparator()

	fmt.Fprint(out, n)
}

func main() {
	run(os.Stdin, os.Stdout)
}
