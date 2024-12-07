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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	h, _, d := readInt(), readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	points := make([][]int, 0)
	for i, s := range ss {
		for j, v := range s {
			if v == '#' {
				continue
			}
			points = append(points, []int{i, j, 0})
		}
	}

	for i, p1 := range points {
		e := 0
		for _, p2 := range points {
			if isIn(p1, p2, d) {
				e++
			}
		}
		points[i] = []int{p1[0], p1[1]}
	}

	ans := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			sum := 0
			p1, p2 := points[i], points[j]
			for _, p3 := range points {
				if isIn(p1, p3, d) {
					sum++
				} else if isIn(p2, p3, d) {
					sum++
				}
			}

			if ans < sum {
				ans = sum
			}
		}
	}
	fmt.Fprint(out, ans)
}

func isIn(a, b []int, d int) bool {
	x := a[0] - b[0]
	y := a[1] - b[1]
	return abs(x)+abs(y) <= d
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
