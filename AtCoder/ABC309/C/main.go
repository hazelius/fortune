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

type med struct {
	a int
	b int
}

type meds []med

func (a meds) Len() int           { return len(a) }
func (a meds) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a meds) Less(i, j int) bool { return a[i].a < a[j].a }

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	abs := make(meds, n)

	num := 0
	for i := range abs {
		abs[i] = med{readInt(), readInt()}
		num += abs[i].b
	}

	sort.Sort(abs)

	days := 1
	for _, ab := range abs {
		if num <= k {
			fmt.Fprint(out, days)
			return
		}
		days = ab.a + 1
		num -= ab.b
	}

	fmt.Fprint(out, days)
}

func main() {
	run(os.Stdin, os.Stdout)
}
