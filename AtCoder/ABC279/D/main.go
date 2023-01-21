// 三分探索
// https://atcoder.jp/contests/abc279/tasks/abc279_d

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	a, b := readInt(), readInt()

	low := 0
	high := a / b
	for low+2 < high {
		mid1 := low + (high-low)/3
		mid2 := low + 2*(high-low)/3
		mid1r := f(mid1, a, b)
		mid2r := f(mid2, a, b)
		// fmt.Printf("%v %.10f %.10f\n", mid1, mid1r, mid1r-mid2r)
		if mid2r < mid1r {
			low = mid1
		} else {
			high = mid2
		}
	}

	ans := math.Min(f(low, a, b), f(low+1, a, b))

	fmt.Fprintf(out, "%.10f", ans)
}

func f(n, a, b int) float64 {
	return float64(b*n) + (float64(a) / math.Sqrt(float64(n+1)))
}

func main() {
	run(os.Stdin, os.Stdout)
}
