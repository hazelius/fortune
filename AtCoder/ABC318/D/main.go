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

	n := readInt()
	ds := make([][]int, n)
	for i := range ds {
		ds[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			ds[i][j] = readInt()
		}
	}

	used := make([]bool, n)
	ans := f(0, ds, used)
	fmt.Fprint(out, ans)
}

func f(c int, ds [][]int, used []bool) int {
	for i, v := range used {
		if v {
			continue
		}
		used[i] = true

		m := c
		r := f(c, ds, used)
		if m < r {
			m = r
		}

		for j := i + 1; j < len(used); j++ {
			if used[j] {
				continue
			}
			used[j] = true
			r := f(c+ds[i][j], ds, used)
			if m < r {
				m = r
			}
			used[j] = false
		}

		used[i] = false
		return m
	}

	return c
}

func main() {
	run(os.Stdin, os.Stdout)
}
