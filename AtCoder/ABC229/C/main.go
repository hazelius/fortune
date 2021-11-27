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

type chs struct {
	a int
	b int
}

type Cheese []chs

func (a Cheese) Len() int           { return len(a) }
func (a Cheese) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Cheese) Less(i, j int) bool { return a[i].a > a[j].a }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, w := readInt(), readInt()
	abs := make(Cheese, n)
	for i := range abs {
		abs[i] = chs{readInt(), readInt()}
	}
	sort.Sort(abs)

	ans := 0
	weight := 0
	for _, ab := range abs {
		if weight+ab.b <= w {
			ans += ab.a * ab.b
			weight += ab.b
		} else {
			w -= weight
			ans += ab.a * w
			return ans
		}
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
