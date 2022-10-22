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

type point struct {
	x int
	y int
}

var memoMap map[point]bool

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, x, y := readInt(), readInt(), readInt()
	as := make([]int, n)
	sum := 0
	for i := range as {
		a := readInt()
		as[i] = a
		sum += a
	}
	memoMap = make(map[point]bool)

	if dfs(point{x: as[0], y: 0}, false, 1, sum-as[0], as, point{x: x, y: y}) {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")
	}
}

func dfs(current point, dirflg bool, cnt, remain int, as []int, target point) bool {
	if cnt == len(as) {
		if current.x == target.x && current.y == target.y {
			return true
		}
		memoMap[current] = true
		return false
	}

	if memoMap[current] {
		return false
	}

	d := abs(current.x-target.x) + abs(current.y-target.y)
	if d > remain {
		return false
	}

	a := as[cnt]
	dirs := [][]int{{0, a}, {0, -a}}
	if dirflg {
		dirs = [][]int{{a, 0}, {-a, 0}}
	}

	cnt++
	remain -= a
	for _, dir := range dirs {
		next := point{x: current.x + dir[0], y: current.y + dir[1]}
		if dfs(next, !dirflg, cnt, remain, as, target) {
			return true
		}
	}

	memoMap[current] = true
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
