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

	a, b, c, d := readInt(), readInt(), readInt(), readInt()
	if a < 0 {
		offset := -((a - 4) / 4) * 4
		a += offset
		c += offset
	}
	if b < 0 {
		offset := -(b + b%2)
		b += offset
		d += offset
	}

	ans := f(c, d) - f(a, d) - f(c, b) + f(a, b)
	fmt.Fprint(out, ans)
}

var ss = [][]int{
	{0, 2, 3, 3},
	{0, 1, 3, 4},
}

func f(a, b int) int {
	ret := (a / 4) * 4 * b
	ret += ((b / 2) + b%2) * ss[0][a%4]
	ret += (b / 2) * ss[1][a%4]
	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
