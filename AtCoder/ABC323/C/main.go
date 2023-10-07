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

func readString() string {
	sc.Scan()
	return sc.Text()
}

type problem struct {
	point int
	idx   int
}

type problems []problem

func (a problems) Len() int           { return len(a) }
func (a problems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a problems) Less(i, j int) bool { return a[i].point > a[j].point }

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	as := make(problems, m)
	for i := range as {
		a := readInt()
		as[i] = problem{point: a, idx: i}
	}

	ss := make([][]bool, n)
	ssum := make([]int, n)
	for i := range ss {
		ss[i] = make([]bool, m)
		ssum[i] = i + 1
		s := readString()
		for j, v := range s {
			if v == 'o' {
				ss[i][j] = true
				ssum[i] += as[j].point
			}
		}
	}

	maxP := 0
	for _, v := range ssum {
		if maxP < v {
			maxP = v
		}
	}

	sort.Sort(as)

	for i, sum := range ssum {
		ans := 0
		for _, v := range as {
			if sum >= maxP {
				break
			}
			if ss[i][v.idx] {
				continue
			}
			sum += v.point
			ans++
		}
		fmt.Fprintln(out, ans)
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
