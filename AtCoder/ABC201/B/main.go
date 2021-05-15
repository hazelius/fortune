package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type mountain struct {
	s string
	t int
}

type SortBy []mountain

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].t < a[j].t }

func run(n int, ms []mountain) string {
	sort.Sort(SortBy(ms))
	ans := ms[len(ms)-2]

	return ans.s
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	ms := make([]mountain, n)
	for i := 0; i < n; i++ {
		ms[i] = mountain{readString(), readInt()}
	}
	fmt.Println(run(n, ms))
}
