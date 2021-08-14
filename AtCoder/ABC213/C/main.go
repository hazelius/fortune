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
	i int
	a int
	b int
}

type SortByA []card

func (a SortByA) Len() int           { return len(a) }
func (a SortByA) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByA) Less(i, j int) bool { return a[i].a < a[j].a }

type SortByB []card

func (a SortByB) Len() int           { return len(a) }
func (a SortByB) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByB) Less(i, j int) bool { return a[i].b < a[j].b }

type SortByI []card

func (a SortByI) Len() int           { return len(a) }
func (a SortByI) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByI) Less(i, j int) bool { return a[i].i < a[j].i }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	_, _, n := readInt(), readInt(), readInt()
	cards := make([]card, n)
	for i := range cards {
		cards[i] = card{i: i, a: readInt(), b: readInt()}
	}

	sort.Sort(SortByA(cards))
	idx, pre := 0, 0
	for i, c := range cards {
		if c.a != pre {
			idx++
		}
		pre = c.a
		cards[i].a = idx
	}

	sort.Sort(SortByB(cards))
	idx, pre = 0, 0
	for i, c := range cards {
		if c.b != pre {
			idx++
		}
		pre = c.b
		cards[i].b = idx
	}

	sort.Sort(SortByI(cards))
	ans := make([]string, n)
	for i, c := range cards {
		ans[i] = fmt.Sprintf("%v %v", c.a, c.b)
	}
	return strings.Join(ans, "\n")
}

func main() {
	fmt.Println(run(os.Stdin))
}
