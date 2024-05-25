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

type kukan struct {
	l int
	r int
}

type kukans []kukan

func (a kukans) Len() int           { return len(a) }
func (a kukans) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a kukans) Less(i, j int) bool { return a[i].l < a[j].l }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	lrs := make(kukans, n)
	for i := range lrs {
		lrs[i] = kukan{readInt(), readInt()}
	}

	sort.Sort(lrs)

	ans := 0
	for i, lr := range lrs {
		r := lr.r
		idx := sort.Search(n, func(j int) bool {
			if i == j {
				return false
			}
			if lrs[j].l > r {
				return true
			}
			return false
		})
		ans += idx - 1 - i
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
