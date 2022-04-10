package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var n, x, y int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, x, y = readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	ans := 0
	start := 0
	for i, a := range as {
		if y <= a && a <= x {
			continue
		}

		if start < i {
			ans += f(as[start:i])
		}
		start = i + 1
	}
	if start < len(as) {
		ans += f(as[start:])
	}

	return ans
}

func f(as []int) int {
	cx, cy := 0, 0
	ret := 0
	rgt := 0

	for i, a := range as {
		if a == x {
			cx++
		}
		if a == y {
			cy++
		}
		for cx > 0 && cy > 0 {
			ret += len(as) - i
			if as[rgt] == x {
				cx--
			}
			if as[rgt] == y {
				cy--
			}
			rgt++
		}
	}
	return ret
}

func main() {
	fmt.Println(run(os.Stdin))
}
