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

	t := readInt()
	for i := 0; i < t; i++ {
		n, d, k := readInt(), readInt(), readInt()-1
		g := GCD(n, d)
		a := n / g
		fmt.Fprintln(out, d*k%n+k/a)
	}
}

func GCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return 0
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return b
}

func main() {
	run(os.Stdin, os.Stdout)
}
