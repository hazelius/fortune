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

	a, x, m := readInt(), readInt(), readInt()
	ans, _ := f(a, x, m)
	fmt.Fprint(out, ans%m)
}

func f(a, x, m int) (int, int) {
	if x == 0 {
		return 0, 1
	}

	s, p := f(a, x/2, m)
	s = (s + s*p) % m
	p = (p * p) % m

	if x%2 > 0 {
		s = (s*a + 1) % m
		p = (p * a) % m
	}

	return s, p
}

func main() {
	run(os.Stdin, os.Stdout)
}
