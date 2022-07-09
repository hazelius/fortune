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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	sx, sy, tx, ty := readInt(), readInt(), readInt(), readInt()
	xyrs := make([][]int, n)
	for i := range xyrs {
		xyrs[i] = []int{readInt(), readInt(), readInt()}
	}
	si := 0
	for i, xyr := range xyrs {
		xl := sx - xyr[0]
		yl := sy - xyr[1]
		if xl*xl+yl*yl == xyr[2]*xyr[2] {
			si = i
			break
		}
	}
	ti := 0
	for i, xyr := range xyrs {
		xl := tx - xyr[0]
		yl := ty - xyr[1]
		if xl*xl+yl*yl == xyr[2]*xyr[2] {
			ti = i
			break
		}
	}
	if si == ti {
		return "Yes"
	}

	used := make([]bool, n)
	used[si] = true

	if f(si, ti, xyrs, used) {
		return "Yes"
	}

	return "No"
}

func f(si, ti int, xyrs [][]int, used []bool) bool {
	x, y, r := xyrs[si][0], xyrs[si][1], xyrs[si][2]
	for i, xyr := range xyrs {
		if used[i] {
			continue
		}
		xl := x - xyr[0]
		yl := y - xyr[1]
		xy2 := xl*xl + yl*yl
		if xy2 > (xyr[2]+r)*(xyr[2]+r) || xy2 < (xyr[2]-r)*(xyr[2]-r) {
			continue
		}
		used[i] = true
		if i == ti {
			return true
		}
		if f(i, ti, xyrs, used) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(run(os.Stdin))
}
