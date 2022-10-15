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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	h, w, rs, cs := readInt(), readInt(), readInt()-1, readInt()-1
	n := readInt()

	rtrees := make(map[int]*rbt.Tree)
	ctrees := make(map[int]*rbt.Tree)

	for i := 0; i < n; i++ {
		r, c := readInt()-1, readInt()-1
		rtree, ok := rtrees[r]
		if !ok {
			rtree = rbt.NewWithIntComparator()
		}
		rtree.Put(c, c)
		rtrees[r] = rtree

		ctree, ok := ctrees[c]
		if !ok {
			ctree = rbt.NewWithIntComparator()
		}
		ctree.Put(r, r)
		ctrees[c] = ctree
	}

	q := readInt()

	for i := 0; i < q; i++ {
		d, l := readString(), readInt()
		nrs, ncs := rs, cs
		switch d {
		case "L":
			ncs -= l
			if ncs < 0 {
				ncs = 0
			}
			tree, ok := rtrees[nrs]
			if ok {
				ob, ok := tree.Floor(cs)
				if ok {
					iob := ob.Key.(int)
					if iob >= ncs {
						ncs = iob + 1
					}
				}
			}

		case "R":
			ncs += l
			if ncs >= w {
				ncs = w - 1
			}
			tree, ok := rtrees[nrs]
			if ok {
				ob, ok := tree.Ceiling(cs)
				if ok {
					iob := ob.Key.(int)
					if iob <= ncs {
						ncs = iob - 1
					}
				}
			}

		case "U":
			nrs -= l
			if nrs < 0 {
				nrs = 0
			}

			tree, ok := ctrees[ncs]
			if ok {
				ob, ok := tree.Floor(rs)
				if ok {
					iob := ob.Key.(int)
					if iob >= nrs {
						nrs = iob + 1
					}
				}
			}
		case "D":
			nrs += l
			if nrs >= h {
				nrs = h - 1
			}
			tree, ok := ctrees[ncs]
			if ok {
				ob, ok := tree.Ceiling(rs)
				if ok {
					iob := ob.Key.(int)
					if iob <= nrs {
						nrs = iob - 1
					}
				}
			}
		}

		rs, cs = nrs, ncs
		fmt.Fprintf(out, "%v %v\n", rs+1, cs+1)
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
