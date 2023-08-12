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

type bet struct {
	i   int
	cnt int
}

type SortBy []bet

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	if a[i].cnt == a[j].cnt {
		return a[i].i < a[j].i
	}
	return a[i].cnt < a[j].cnt
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	cmap := make(map[int][]bet)

	for i := 0; i < n; i++ {
		c := readInt()
		b := bet{i + 1, c}
		for i := 0; i < c; i++ {
			a := readInt()
			cmap[a] = append(cmap[a], b)
		}
	}
	x := readInt()

	bs, ok := cmap[x]
	if !ok {
		fmt.Fprint(out, 0)
		return
	}
	sort.Sort(SortBy(bs))
	ans := make([]int, 0)
	cnt := bs[0].cnt
	for _, b := range bs {
		if b.cnt != cnt {
			break
		}
		ans = append(ans, b.i)
	}

	fmt.Fprintln(out, len(ans))
	ansstr := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, ansstr[1:len(ansstr)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
