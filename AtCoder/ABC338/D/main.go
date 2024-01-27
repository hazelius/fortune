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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	xs := make([]int, m)
	for i := range xs {
		xs[i] = readInt() - 1
	}

	costs := make([]int, n+1)
	for i := 0; i < m-1; i++ {
		a, b := xs[i], xs[i+1]
		d := dist(b, a, n)
		costs[a] += d
		costs[b] -= d
		if a >= b {
			costs[n] -= d
			costs[0] += d
		}

		a, b = xs[i+1], xs[i]
		d = dist(b, a, n)
		costs[a] += d
		costs[b] -= d
		if a >= b {
			costs[n] -= d
			costs[0] += d
		}
	}

	minl := -1
	for i := 0; i < n; i++ {
		costs[i+1] += costs[i]
		if minl < 0 {
			minl = costs[i]
		} else {
			minl = min(minl, costs[i])
		}
	}

	fmt.Fprint(out, minl)
}

func dist(a, b, n int) int {
	if a <= b {
		return b - a
	}
	return b + n - a
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
