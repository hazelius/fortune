package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var memo map[int]float64

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, a, x, y := readInt(), readInt(), readInt(), readInt()
	memo = make(map[int]float64)
	ans := f(n, a, float64(x), float64(y))

	fmt.Fprintf(out, "%.7f", ans)
}

func f(n, a int, x, y float64) float64 {
	if n == 0 {
		return 0
	}
	if r, ok := memo[n]; ok {
		return r
	}

	res := x + f(n/a, a, x, y)

	dice := 0.0
	for i := 2; i <= 6; i++ {
		dice += f(n/i, a, x, y)
	}
	dice /= 5
	dice += y * 6 / 5
	if res > dice {
		res = dice
	}
	memo[n] = res
	return res
}

func main() {
	run(os.Stdin, os.Stdout)
}
