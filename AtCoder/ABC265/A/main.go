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

	x, y, n := readInt(), readInt(), readInt()
	xs := x * n
	ys := (n/3)*y + (n%3)*x
	if xs < ys {
		fmt.Fprint(out, xs)
		return
	}
	fmt.Fprint(out, ys)
}

func main() {
	run(os.Stdin, os.Stdout)
}
