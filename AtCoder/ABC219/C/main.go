package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}
func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type Dict [][]string

func (a Dict) Len() int           { return len(a) }
func (a Dict) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Dict) Less(i, j int) bool { return a[i][0] < a[j][0] }

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	x := readString()
	n := readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	mp := make(map[rune]rune)
	for i, c := range x {
		mp[c] = rune(i)
	}

	mss := make([][]string, n)
	for i, s := range ss {
		ds := make([]rune, len(s))
		for j, c := range s {
			ds[j] = mp[c]
		}
		mss[i] = []string{string(ds), s}
	}

	sort.Sort(Dict(mss))

	ans := make([]string, n)
	for i, s := range mss {
		ans[i] = s[1]
	}
	return strings.Join(ans, " ")
}

func main() {
	fmt.Println(run(os.Stdin))
}
