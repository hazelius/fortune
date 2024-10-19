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

func readFloat() float64 {
	sc.Scan()
	i, _ := strconv.ParseFloat(sc.Text(), 64)
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	p := []float64{0, 0}
	ans := 0.0
	for i := 0; i < n; i++ {
		next := []float64{readFloat(), readFloat()}
		ans += dist(p, next)
		p = next
	}
	ans += dist(p, []float64{0, 0})

	fmt.Fprintf(out, "%.12f", ans)
}

func dist(a, b []float64) float64 {
	x := a[0] - b[0]
	y := a[1] - b[1]
	return math.Sqrt(x*x + y*y)
}

func main() {
	run(os.Stdin, os.Stdout)
}
