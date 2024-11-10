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
	a := n / 100
	b := (n % 100) / 10
	c := n % 10

	fmt.Fprintf(out, "%d%d%d %d%d%d", b, c, a, c, a, b)
}

func main() {
	run(os.Stdin, os.Stdout)
}
