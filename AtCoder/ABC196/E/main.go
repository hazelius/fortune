package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type k struct {
	v int
	b bool
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(fs [][]int, xs []int) []int {
	low, high, z := minimize(fs)
	fmt.Printf("l:%v, h:%v, z:%v\n", low, high, z)
	ret := make([]int, len(xs))
	for i, x := range xs {
		x += z
		if high.b {
			x = max(high.v, x)
		}
		if low.b {
			x = min(low.v, x)
		}
		ret[i] = x
	}
	return ret
}

func minimize(fs [][]int) (k, k, int) {
	var low, high k
	var z int
	for _, f := range fs {
		a := f[0]
		t := f[1]
		switch t {
		case 1:
			z += a
			if low.b {
				low.v += a
			}
			if high.b {
				high.v += a
			}
		case 2:
			if !high.b {
				high = k{v: a, b: true}
			} else if high.v < a {
				high.v = a
			}
			if low.b && low.v < a {
				low.v = a
			}
		case 3:
			if high.b && high.v > a {
				high.v = a
			}
			if !low.b {
				low = k{v: a, b: true}
			} else if low.v > a {
				low.v = a
			}
		}
	}
	return low, high, z
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	fs := make([][]int, n)
	for i := range fs {
		fs[i] = make([]int, 2)
		fs[i][0], fs[i][1] = readInt(), readInt()
	}
	q := readInt()
	xs := make([]int, q)
	for i := range xs {
		xs[i] = readInt()
	}

	for _, r := range run(fs, xs) {
		fmt.Println(r)
	}
}
