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

	n, q := readInt(), readInt()
	ls := make([]int, n)
	for i := 1; i < n; i++ {
		ls[i] = readInt()
	}
	rs := make([]int, n)
	for i := 0; i < n-1; i++ {
		rs[i] = readInt()
	}

	lsMin := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			lsMin[i] = ls[i]
			continue
		}
		lsMin[i] = min(lsMin[i+1], ls[i])
	}

	rsMin := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			rsMin[i] = rs[i]
			continue
		}
		rsMin[i] = min(rsMin[i-1], rs[i])
	}

	for i := 0; i < q; i++ {
		s, t := readInt()-1, readInt()-1
		if s < t {
			cost := rs[s]
			if s > 0 {
				costl := ls[s] + rsMin[s]
				cost = min(cost, costl)
			}
			fmt.Fprintln(out, cost)
		} else {
			cost := ls[s]
			if s < n-1 {
				costr := rs[s] + lsMin[s]
				cost = min(cost, costr)
			}
			fmt.Fprintln(out, cost)
		}
	}
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
