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

type player struct {
	idx   int
	point int
}

type Players []player

func (a Players) Len() int      { return len(a) }
func (a Players) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Players) Less(i, j int) bool {
	if a[i].point == a[j].point {
		return a[i].idx < a[j].idx
	}
	return a[i].point > a[j].point
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make(Players, n)
	for i := range as {
		as[i].idx = i
		s := readString()
		for _, v := range s {
			if v == 'o' {
				as[i].point++
			}
		}
	}

	sort.Sort(as)

	for _, p := range as {
		fmt.Fprintf(out, "%v ", p.idx+1)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
