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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, a, b := readInt(), readInt(), readInt()
	sum := 0
	for i := a; i <= n; i += a {
		sum += i
	}
	for i := b; i <= n; i += b {
		sum += i
	}
	ablcm := LCM(a, b)
	for i := ablcm; i <= n; i += ablcm {
		sum -= i
	}

	ans := ((n + 1) * n) / 2
	return ans - sum
}

func LCM(a, b int) int {
	return a / GCD(a, b) * b
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
	fmt.Println(run(os.Stdin))
}
