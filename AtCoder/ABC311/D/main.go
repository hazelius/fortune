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

	n, _ := readInt(), readInt()
	as := make([]string, n)
	for i := range as {
		as[i] = readString()
	}

	ansmap := make(map[point]bool)
	used := make(map[point]bool)
	que := []point{{1, 1}}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		if used[p] {
			continue
		}

		ansmap[p] = true
		used[p] = true

		for _, dir := range dirs {
			next := point{p.x + dir[0], p.y + dir[1]}
			for as[next.x][next.y] != '#' {
				ansmap[next] = true
				next.x += dir[0]
				next.y += dir[1]
			}
			next.x -= dir[0]
			next.y -= dir[1]

			if !used[next] {
				que = append(que, next)
			}
		}
	}

	fmt.Fprint(out, len(ansmap))
}

func main() {
	run(os.Stdin, os.Stdout)
}
