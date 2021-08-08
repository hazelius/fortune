package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

type card struct {
	i   int
	val int
}

type SortBy []card

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].val < a[j].val }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	_, _, n := readInt(), readInt(), readInt()
	as, bs := make([]card, n), make([]card, n)
	for i := range as {
		as[i] = card{i: i, val: readInt()}
		bs[i] = card{i: i, val: readInt()}
	}

	sort.Sort(SortBy(as))
	sort.Sort(SortBy(bs))
	ansa := make([]int, n)
	ansb := make([]int, n)
	idxa, idxb := 1, 1
	for i := 0; i < n; i++ {
		a, b := as[i], bs[i]

		if i > 0 {
			if as[i-1].val != a.val {
				idxa++
			}
			if bs[i-1].val != b.val {
				idxb++
			}
		}

		ansa[a.i] = idxa
		ansb[b.i] = idxb
	}

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = fmt.Sprintf("%v %v", ansa[i], ansb[i])
	}
	return strings.Join(ans, "\n")
}

func main() {
	fmt.Println(run(os.Stdin))
}
