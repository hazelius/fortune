package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type mem struct {
	index int
	a     int
	b     int
}

type mems []mem

func (a mems) Len() int      { return len(a) }
func (a mems) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a mems) Less(i, j int) bool {
	v1 := a[i].a * (a[j].a + a[j].b)
	v2 := a[j].a * (a[i].a + a[i].b)
	if v1 == v2 {
		return a[i].index < a[j].index
	}
	return v1 > v2
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ab := make(mems, n)
	for i := range ab {
		a, b := readInt(), readInt()
		ab[i] = mem{index: i + 1, a: a, b: b}
	}
	sort.Sort(ab)
	for _, m := range ab {
		fmt.Fprintf(out, "%v ", m.index)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
