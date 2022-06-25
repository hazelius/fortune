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

func readString() string {
	sc.Scan()
	return sc.Text()
}

type SortBy [][]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i][0] < a[j][0] }

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	s := readString()
	ws := make([][]int, n)
	for i := range ws {
		f := 0
		if s[i] == '1' {
			f = 1
		}
		ws[i] = []int{readInt(), f}
	}

	sort.Sort(SortBy(ws))
	hit0, hit1 := 0, 0
	for _, v := range ws {
		if v[1] == 1 {
			hit1++
		}
	}

	ans := hit1
	for i, v := range ws {
		if v[1] == 0 {
			hit0++
		} else {
			hit1--
		}
		if i == n-1 || v[0] != ws[i+1][0] {
			if ans < hit0+hit1 {
				ans = hit0 + hit1
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
