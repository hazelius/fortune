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

var sn = "snuke"
var h int
var w int

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w = readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	used := make(map[point]bool)
	if f(0, point{0, 0}, ss, used) {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")
	}

}

func f(i int, p point, ss []string, used map[point]bool) bool {
	if used[p] {
		return false
	}
	if sn[i] != ss[p.x][p.y] {
		return false
	}

	if p.x == h-1 && p.y == w-1 {
		return true
	}
	used[p] = true
	i = (i + 1) % len(sn)

	if p.x > 0 {
		if f(i, point{p.x - 1, p.y}, ss, used) {
			return true
		}
	}
	if p.x < h-1 {
		if f(i, point{p.x + 1, p.y}, ss, used) {
			return true
		}
	}
	if p.y > 0 {
		if f(i, point{p.x, p.y - 1}, ss, used) {
			return true
		}
	}
	if p.y < w-1 {
		if f(i, point{p.x, p.y + 1}, ss, used) {
			return true
		}
	}

	return false
}

func main() {
	run(os.Stdin, os.Stdout)
}
