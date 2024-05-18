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

type card struct {
	idx int
	a   int
	c   int
}
type cards []card

func (a cards) Len() int           { return len(a) }
func (a cards) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a cards) Less(i, j int) bool { return a[i].c < a[j].c }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	cs := make(cards, n)
	for i := range cs {
		cs[i] = card{i + 1, readInt(), readInt()}
	}
	sort.Sort(cs)

	ans := make([]int, 0)
	max := 0
	for _, c := range cs {
		if c.a > max {
			max = c.a
			ans = append(ans, c.idx)
		}
	}
	fmt.Fprintln(out, len(ans))
	sort.Ints(ans)
	anss := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, anss[1:len(anss)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
