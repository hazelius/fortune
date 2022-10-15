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

	n := readInt()
	fmt.Fprint(out, f(n))
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	return x * f(x-1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
