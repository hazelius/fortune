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

type katamuki struct {
	x int
	y int
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	xys := make([][]int, n)
	for i := range xys {
		xys[i] = []int{readInt(), readInt()}
	}

	kmap := make(map[katamuki]bool)
	for i := 0; i < n-1; i++ {
		xi, yi := xys[i][0], xys[i][1]
		for j := i + 1; j < n; j++ {
			x := xys[j][0] - xi
			y := xys[j][1] - yi

			g := GCD(x, y)
			if g == 0 {
				if x == 0 {
					kmap[katamuki{0, 1}] = true
					kmap[katamuki{0, -1}] = true
				} else {
					kmap[katamuki{1, 0}] = true
					kmap[katamuki{-1, 0}] = true
				}
			} else {
				kmap[katamuki{x / g, y / g}] = true
				kmap[katamuki{(x / g) * -1, (y / g) * -1}] = true
			}
		}
	}

	// kary := make([]katamuki, 0, len(kmap))
	// for k := range kmap {
	// 	flg := true
	// 	for _, l := range kary {
	// 		if k.x*l.y == k.y*l.x {
	// 			flg = false
	// 			break
	// 		}
	// 	}
	// 	if flg {
	// 		kary = append(kary, k)
	// 	}
	// }

	return len(kmap)
}

func GCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return 0
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return b
}

func main() {
	fmt.Println(run(os.Stdin))
}
