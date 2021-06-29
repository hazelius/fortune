// D - Congruence Points
// https://atcoder.jp/contests/abc207/tasks/abc207_d
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

type point struct {
	x int
	y int
}

type tri struct {
	x int
	y int
	g int
}

type SortBy []tri

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	if a[i].x == a[j].x {
		if a[i].y == a[j].y {
			return a[i].g < a[j].g
		}
		return a[i].y < a[j].y
	}
	return a[i].x < a[j].x
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	abs := make([]point, n)
	for i := range abs {
		abs[i] = point{readInt(), readInt()}
	}
	cds := make([]point, n)
	for i := range cds {
		cds[i] = point{readInt(), readInt()}
	}

	if n == 1 {
		return "Yes"
	}
	if n == 2 {
		if dist(abs[0], abs[1]) == dist(cds[0], cds[1]) {
			return "Yes"
		}
		return "No"
	}

	abt := make([]tri, n)
	for i := range abt {
		abt[i] = tri{
			x: dist(abs[0], abs[i]),
			y: dist(abs[1], abs[i]),
			g: g(abs[0], abs[1], abs[i]),
		}
	}
	sort.Sort(SortBy(abt))

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			cdt := make([]tri, n)
			for k := range cdt {
				cdt[k] = tri{
					x: dist(cds[i], cds[k]),
					y: dist(cds[j], cds[k]),
					g: g(cds[i], cds[j], cds[k]),
				}
			}
			sort.Sort(SortBy(cdt))

			if eq(abt, cdt) {
				return "Yes"
			}
		}
	}
	return "No"
}

func dist(a, b point) int {
	x := a.x - b.x
	y := a.y - b.y
	return x*x + y*y
}

func g(a, b, c point) int {
	return (a.x-c.x)*(b.y-c.y) - (a.y-c.y)*(b.x-c.x)
}

func eq(a, b []tri) bool {
	for i := range a {
		if !(a[i].x == b[i].x && a[i].y == b[i].y && a[i].g == b[i].g) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(run(os.Stdin))
}
