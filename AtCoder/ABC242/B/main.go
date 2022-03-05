package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

type SortBy []rune

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()
	ans := SortBy(s)
	sort.Sort(ans)
	return string(ans)
}

func main() {
	fmt.Println(run(os.Stdin))
}
