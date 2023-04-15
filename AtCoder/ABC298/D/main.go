package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var mod = 998244353

func Mod(a int) int {
	a %= mod
	if a < 0 {
		a += mod
	}
	return a
}

func modPow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 == 1 {
			p = Mod(p * a)
		}
		a = Mod(a * a)
		b >>= 1
	}

	return p
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	q := readInt()
	s := 1
	ss := make([]int, 1, q)
	ss[0] = 1
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x := readInt()
			s = Mod(10*s + x)
			ss = append(ss, x)
		case 2:
			ms := modPow(10, len(ss)-1)
			ms = Mod(ms * ss[0])
			s = Mod(s - ms)
			ss = ss[1:]
		case 3:
			fmt.Fprintln(out, s)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
