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

type item struct {
	price int
	f     map[int]bool
}

type Items []item

func (a Items) Len() int      { return len(a) }
func (a Items) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Items) Less(i, j int) bool {
	if a[i].price == a[j].price {
		return len(a[i].f) > len(a[j].f)
	}
	return a[i].price < a[j].price
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, _ := readInt(), readInt()
	its := make(Items, n)
	for i := range its {
		its[i].price = readInt()
		c := readInt()
		fm := make(map[int]bool)
		for j := 0; j < c; j++ {
			f := readInt()
			fm[f] = true
		}
		its[i].f = fm
	}

	sort.Sort(its)

	for i, it := range its {
		for j := i + 1; j < n; j++ {
			it2 := its[j]
			if it.price == it2.price && len(it.f) <= len(it2.f) {
				continue
			}
			flg := true
			for f := range it2.f {
				if !it.f[f] {
					flg = false
					break
				}
			}
			if flg {
				fmt.Fprint(out, "Yes")
				return
			}
		}
	}
	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
