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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, a, b := readInt(), readInt(), readInt()
	ds := make([]int, n)
	for i := range ds {
		ds[i] = readInt() % (a + b)
	}
	sort.Ints(ds)

	maxday, minday := 0, a+b
	for _, d := range ds {
		if maxday < d {
			maxday = d
		}
		if minday > d {
			minday = d
		}
	}

	if maxday-minday < a {
		fmt.Fprint(out, "Yes")
		return
	}

	for i := 0; i < n-1; i++ {
		v1, v2 := ds[i], ds[i+1]

		if (a+b)-(v2-v1) < a {
			fmt.Fprint(out, "Yes")
			return
		}
	}

	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
