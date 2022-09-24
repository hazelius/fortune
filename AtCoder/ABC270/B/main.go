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

	x, y, z := readInt(), readInt(), readInt()
	if x < 0 {
		x *= -1
		y *= -1
		z *= -1
	}

	if y < 0 || x < y {
		fmt.Fprint(out, x)
		return
	}

	if y < z {
		fmt.Fprint(out, -1)
		return
	}

	ans := abs(z) + x - z

	fmt.Fprint(out, ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
