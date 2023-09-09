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

type card struct {
	idx int
	num int
}

type cards []card

func (a cards) Len() int           { return len(a) }
func (a cards) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a cards) Less(i, j int) bool { return a[i].idx < a[j].idx }

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	cs := make([]int, 9)
	for i := range cs {
		cs[i] = readInt()
	}
	cardIdx := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	dirs := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	targets := make([][]int, 0, len(dirs))
	for _, dir := range dirs {
		if cs[dir[0]] != cs[dir[1]] && cs[dir[0]] != cs[dir[2]] && cs[dir[1]] != cs[dir[2]] {
			continue
		}
		targets = append(targets, dir)
	}

	all := 0.0
	gakkari := 0.0
	for {
		all++

		for _, d := range targets {
			cs := cards{
				{cardIdx[d[0]], cs[d[0]]},
				{cardIdx[d[1]], cs[d[1]]},
				{cardIdx[d[2]], cs[d[2]]},
			}
			sort.Sort(cs)
			if cs[0].num == cs[1].num {
				gakkari++
				break
			}
		}

		if !NextPermutation(sort.IntSlice(cardIdx)) {
			break
		}
	}

	fmt.Fprintf(out, "%.12f", (all-gakkari)/all)
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	run(os.Stdin, os.Stdout)
}
