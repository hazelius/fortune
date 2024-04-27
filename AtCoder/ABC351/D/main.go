package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var used map[point]bool
var dir = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type point struct {
	x int
	y int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	used = make(map[point]bool)
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			p := point{i, j}
			if used[p] {
				continue
			}
			cntused := make(map[point]bool)
			ret := dfs(point{i, j}, ss, cntused)
			if ans < ret {
				ans = ret
			}
		}
	}
	fmt.Fprint(out, ans)
}

func dfs(p point, ss []string, cntused map[point]bool) int {
	if ss[p.x][p.y] == '#' {
		return 0
	}
	if cntused[p] {
		return 0
	}
	used[p] = true
	cntused[p] = true

	newps := make([]point, 0)
	for _, d := range dir {
		newp := point{p.x + d[0], p.y + d[1]}
		if newp.x < 0 || newp.x >= len(ss) || newp.y < 0 || newp.y >= len(ss[0]) {
			continue
		}

		if ss[newp.x][newp.y] == '#' {
			return 1
		}
		if !cntused[newp] {
			newps = append(newps, newp)
		}
	}

	ret := 1
	for _, newp := range newps {
		ret += dfs(newp, ss, cntused)
	}

	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
