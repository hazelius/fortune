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

type Line struct {
	a int
	b int
}

type Lines []Line

func (a Lines) Len() int      { return len(a) }
func (a Lines) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Lines) Less(i, j int) bool {
	if a[i].a < a[j].a {
		return true
	}
	if a[i].a > a[j].a {
		return false
	}
	return a[i].b < a[j].b
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	_, m := readInt(), readInt()
	abs := make(Lines, m)
	for i := range abs {
		abs[i] = Line{readInt(), readInt()}
	}

	sort.Sort(abs)

	ans := 0
	invs := make([]Line, 0, m)
	cnts := make([]int, 0, m)
	for i := 0; i < m; i++ {
		t := abs[i]
		flg := false
		for idx, l := range invs {
			if l.a < t.a && l.b < t.b {
				invs[idx] = t
				cnts[idx]++
				flg = true
				break
			}
		}
		if !flg {
			invs = append(invs, t)
			cnts = append(cnts, 1)
		}
	}

	for _, v := range cnts {
		if v > ans {
			ans = v
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
