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

	a, b := readInt(), readInt()

	fmt.Fprintf(out, "%.03f", float64(b)/float64(a))
}

func main() {
	run(os.Stdin, os.Stdout)
}
