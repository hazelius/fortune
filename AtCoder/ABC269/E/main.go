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
	ans := make([]int, 2)
	for i := range ans {
		low, high := 1, n+1
		for low+1 < high {
			mid := low + (high-low)/2
			a, b, c, d := low, mid-1, 1, n
			if i > 0 {
				a, b, c, d = 1, n, low, mid-1
			}
			fmt.Fprintf(out, "? %v %v %v %v\n", a, b, c, d)
			ret := readInt()
			if ret == (mid - low) {
				low = mid
			} else {
				high = mid
			}
		}
		ans[i] = low
	}
	fmt.Fprintf(out, "! %v %v\n", ans[0], ans[1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
