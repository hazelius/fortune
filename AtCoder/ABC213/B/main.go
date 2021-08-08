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

type player struct {
	i     int
	score int
}

type SortBy []player

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].score > a[j].score }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]player, n)
	for i := range as {
		as[i] = player{i: i + 1, score: readInt()}
	}
	sort.Sort(SortBy(as))
	return as[1].i
}

func main() {
	fmt.Println(run(os.Stdin))
}
