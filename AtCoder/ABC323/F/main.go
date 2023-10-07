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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	a := point{readInt(), readInt()}
	b := point{readInt(), readInt()}
	x := point{readInt(), readInt()}

	start, ans := startPosition(a, b, x)
	ans += distance(b, x)
	if start.x != x.x && start.y != x.y {
		ans += 2
	}

	fmt.Fprint(out, ans)
}

func startPosition(a, b, x point) (point, int) {
	ret1 := b
	if x.y < b.y {
		ret1.y++
	} else if x.y > b.y {
		ret1.y--
	}

	ret2 := b
	if x.x < b.x {
		ret2.x++
	} else if x.x > b.x {
		ret2.x--
	}

	if ret1.x == b.x && ret1.y == b.y {
		dis := distance(a, ret2)
		if isDisurb(a, b, ret2) {
			dis += 2
		}
		return ret2, dis
	}
	if ret2.x == b.x && ret2.y == b.y {
		dis := distance(a, ret1)
		if isDisurb(a, b, ret1) {
			dis += 2
		}
		return ret1, dis
	}

	dis1 := distance(a, ret1)
	dis2 := distance(a, ret2)
	if isDisurb(a, b, ret1) {
		dis1 += 2
	}
	if isDisurb(a, b, ret2) {
		dis2 += 2
	}

	if dis1 < dis2 {
		return ret1, dis1
	}
	return ret2, dis2
}

func isDisurb(a, b, dis point) bool {
	if a.x == dis.x && dis.x == b.x {
		if a.y < b.y && b.y < dis.y || a.y > b.y && b.y > dis.y {
			return true
		}
	}
	if a.y == dis.y && dis.y == b.y {
		if a.x < b.x && b.x < dis.x || a.x > b.x && b.x > dis.x {
			return true
		}
	}
	return false
}

func distance(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
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
