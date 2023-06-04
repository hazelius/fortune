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

	n, d := readInt(), readInt()
	xys := make([][]int, n)
	for i := range xys {
		xys[i] = []int{readInt(), readInt()}
	}

	d *= d

	q := []int{0}
	used := make([]bool, n)
	used[0] = true
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for i, xy := range xys {
			if used[i] {
				continue
			}
			if f(xys[p], xy, d) {
				used[i] = true
				q = append(q, i)
			}
		}
	}

	for _, b := range used {
		if b {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func f(a, b []int, d int) bool {
	x := a[0] - b[0]
	y := a[1] - b[1]
	return (x*x + y*y) <= d
}

func main() {
	run(os.Stdin, os.Stdout)
}
