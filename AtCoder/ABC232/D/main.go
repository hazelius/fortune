package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

var (
	h      int
	w      int
	m      int
	passed [][]bool
)

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	h, w = readInt(), readInt()
	cs := make([]string, h)
	for i := range cs {
		cs[i] = readString()
	}

	m = 0
	passed = make([][]bool, h)
	for i := range passed {
		passed[i] = make([]bool, w)
	}
	dfs(0, 0, 0, cs)

	return m
}

func dfs(i, j, cnt int, cs []string) {
	if cs[i][j] == '#' {
		return
	}

	cnt++
	if cnt > m {
		m = cnt
	}
	passed[i][j] = true

	if i+1 < h {
		if !passed[i+1][j] {
			dfs(i+1, j, cnt, cs)
		}
	}
	if j+1 < w {
		if !passed[i][j+1] {
			dfs(i, j+1, cnt, cs)
		}
	}
}

func main() {
	fmt.Println(run(os.Stdin))
}
