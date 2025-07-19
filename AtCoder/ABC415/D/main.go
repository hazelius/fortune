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

type Bin struct {
	a int
	b int
}

func (ab *Bin) sabun() int {
	return ab.a - ab.b
}

type Bins []Bin

func (a Bins) Len() int      { return len(a) }
func (a Bins) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Bins) Less(i, j int) bool {
	if a[i].sabun() == a[j].sabun() {
		return a[i].a < a[j].a
	}
	return a[i].sabun() < a[j].sabun()
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	abs := make(Bins, m)
	for i := range abs {
		abs[i] = Bin{readInt(), readInt()}
	}
	sort.Sort(abs)

	ans := 0
	for i := 0; i < m; i++ {
		ab := abs[i]
		if ab.a > n {
			continue
		}
		cnt := (n - ab.a) / ab.sabun()
		cnt++
		ans += cnt
		n -= ab.sabun() * cnt
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
