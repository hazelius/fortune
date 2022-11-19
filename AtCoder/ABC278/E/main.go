package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type Range struct {
	minh int
	maxh int
	minw int
	maxw int
}

func (r *Range) addMaxh(a int) {
	r.maxh = max(r.maxh, a)
}
func (r *Range) addMinh(a int) {
	r.minh = min(r.minh, a)
}
func (r *Range) addMaxw(a int) {
	r.maxw = max(r.maxw, a)
}
func (r *Range) addMinw(a int) {
	r.minw = min(r.minw, a)
}
func (r *Range) valid(h, w, hm, wm int) bool {
	if r.maxh > hm {
		return true
	}
	if r.maxw > wm {
		return true
	}
	if r.minh < h {
		return true
	}
	if r.minw < w {
		return true
	}
	return false
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	H, W, n, h, w := readInt(), readInt(), readInt(), readInt(), readInt()

	ranges := make([]Range, n+1)
	for i := range ranges {
		ranges[i].minh = H
		ranges[i].minw = W
	}

	for ih := 0; ih < H; ih++ {
		for iw := 0; iw < W; iw++ {
			a := readInt()
			ranges[a].addMaxh(ih)
			ranges[a].addMaxw(iw)
			ranges[a].addMinh(ih)
			ranges[a].addMinw(iw)
		}
	}

	for ih := 0; ih <= H-h; ih++ {
		for iw := 0; iw <= W-w; iw++ {
			ans := n
			for i := 1; i <= n; i++ {
				r := ranges[i]
				// fmt.Printf("%v %v %v %v %v %v\n", r, ih, iw, ih+h-1, iw+w-1, r.valid(ih, iw, ih+h-1, iw+w-1))
				if !r.valid(ih, iw, ih+h-1, iw+w-1) {
					ans--
				}
			}
			fmt.Fprintf(out, "%v ", ans)
		}
		fmt.Fprintln(out, "")
	}

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
