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

type SortBy [][]int

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	}
	return a[i][0] < a[j][0]
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	lrs := make([][]int, n)
	for i := range lrs {
		lrs[i] = []int{readInt(), readInt()}
	}

	sort.Sort(SortBy(lrs))
	ans := make([][]int, 0, n)
	last := []int{}
	for _, lr := range lrs {
		if len(last) == 0 {
			last = lr
			continue
		}
		if last[1] >= lr[0] {
			if last[1] < lr[1] {
				last[1] = lr[1]
			}
		} else {
			ans = append(ans, last)
			last = lr
		}
	}
	ans = append(ans, last)

	ansarr := make([]string, len(ans))
	for i, v := range ans {
		s := fmt.Sprintf("%v", v)
		ansarr[i] = s[1 : len(s)-1]
	}
	ansstr := fmt.Sprintf("%v", ansarr)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
