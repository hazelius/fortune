package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type FenwickTree []int

func NewFenwickTree(n int) FenwickTree {
	return make(FenwickTree, n)
}
func (f FenwickTree) Add(p, x int) {
	for p++; p <= len(f); p += p & -p {
		f[p-1] += x
	}
}
func (f FenwickTree) Sum(l, r int) int {
	return f.sum(r) - f.sum(l)
}
func (f FenwickTree) sum(r int) int {
	s := 0
	for ; r > 0; r -= r & -r {
		s += f[r-1]
	}
	return s
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
	f := NewFenwickTree(n)
	for i := 0; i < n; i++ {
		a := readInt()
		f.Add(i, a)
		f.Add(i+1, -a)
	}

	for i := 0; i < m; i++ {
		b := readInt()
		a := f.Sum(0, b+1)
		f.Add(b, -a)
		f.Add(b+1, a)

		f.Add(0, a/n)

		l := a % n

		f.Add(b+1, 1)
		if b+l >= n {
			f.Add(0, 1)
			f.Add(b+l-n+1, -1)
		} else {
			f.Add(b+1+l, -1)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(out, f.Sum(0, i+1), " ")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
