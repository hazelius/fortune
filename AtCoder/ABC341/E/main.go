package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	s := readString()

	ft := NewFenwickTree(n)
	for i := range s {
		if i == 0 {
			continue
		}
		if s[i] == s[i-1] {
			ft.Add(i, 1)
		}
	}

	for i := 0; i < q; i++ {
		que, l, r := readInt(), readInt()-1, readInt()-1
		switch que {
		case 1:
			if ft.Sum(l, l+1) == 0 {
				ft.Add(l, 1)
			} else {
				ft.Add(l, -1)
			}
			if r < n-1 {
				if ft.Sum(r+1, r+2) == 0 {
					ft.Add(r+1, 1)
				} else {
					ft.Add(r+1, -1)
				}
			}
		case 2:
			tmp := ft.Sum(l+1, r+1)
			if tmp == 0 {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
