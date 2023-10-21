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

type point struct {
	x int
	y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, _ := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	used := make(map[point]bool)
	ans := 0
	for i, s := range ss {
		for j, c := range s {
			p := point{i, j}
			if c == '#' && !used[p] {
				ans++
				f(p, ss, used)
			}
		}
	}

	fmt.Fprint(out, ans)
}

func f(p point, ss []string, used map[point]bool) {
	dir := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	que := []point{p}
	for len(que) > 0 {
		q := que[0]
		que = que[1:]
		if used[q] {
			continue
		}

		used[q] = true
		for _, d := range dir {
			next := point{q.x + d[0], q.y + d[1]}
			if next.x < 0 || next.y < 0 || next.x >= len(ss) || next.y >= len(ss[0]) {
				continue
			}
			if ss[next.x][next.y] != '#' {
				continue
			}
			if used[next] {
				continue
			}
			que = append(que, next)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
