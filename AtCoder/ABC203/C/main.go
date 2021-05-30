package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type SortBy [][]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i][0] < a[j][0] }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n, k int, abs [][]int) int {
	sort.Sort(SortBy(abs))
	ans := k
	for _, ab := range abs {
		a, b := ab[0], ab[1]
		if a <= ans {
			ans += b
		} else {
			break
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	abs := make([][]int, n)
	for i := range abs {
		abs[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(n, k, abs))
}
