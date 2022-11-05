package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	i int
	j int
}

var used map[point]bool

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, _ := readInt(), readInt()
	hws := make([]string, h)
	for i := range hws {
		hws[i] = readString()
	}

	ans := "No"
	used = make(map[point]bool)
	for i, s := range hws {
		idx := strings.Index(s, "S")
		if idx >= 0 {
			start := point{i: i, j: idx}
			if dfs(start, start, -1, hws) {
				ans = "Yes"
			}
			break
		}
	}

	fmt.Fprint(out, ans)
}

func dfs(start, currnt point, cnt int, hws []string) bool {
	used[currnt] = true
	cnt++

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	validPoint := make([]point, 0)
	for _, dir := range dirs {
		next := point{i: currnt.i + dir[0], j: currnt.j + dir[1]}
		if next.i < 0 || next.j < 0 || next.i >= len(hws) || next.j >= len(hws[0]) {
			continue
		}
		if hws[next.i][next.j] == '#' {
			continue
		}
		if next.i == start.i && next.j == start.j {
			if cnt == 1 {
				continue
			}
			return true
		}
		if used[next] {
			continue
		}
		validPoint = append(validPoint, next)
	}

	for _, p := range validPoint {
		if dfs(start, p, cnt, hws) {
			return true
		}
	}

	return false
}

func main() {
	run(os.Stdin, os.Stdout)
}
