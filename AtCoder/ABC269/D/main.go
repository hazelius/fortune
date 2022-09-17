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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	bmap := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		x, y := readInt(), readInt()
		if bmap[x] == nil {
			bmap[x] = make(map[int]bool)
		}
		bmap[x][y] = true
	}

	ans := 0
	used := make(map[int]map[int]bool)
	for x, ymap := range bmap {
		for y := range ymap {
			if f(x, y, bmap, used) {
				ans++
			}
		}
	}

	fmt.Fprint(out, ans)
}

func f(x, y int, bmap, used map[int]map[int]bool) bool {
	if used[x][y] {
		return false
	}
	if used[x] == nil {
		used[x] = make(map[int]bool)
	}
	used[x][y] = true

	dir := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
	}
	for _, d := range dir {
		xnew := x + d[0]
		ynew := y + d[1]
		if bmap[xnew][ynew] {
			f(xnew, ynew, bmap, used)
		}
	}

	return true
}

func main() {
	run(os.Stdin, os.Stdout)
}
