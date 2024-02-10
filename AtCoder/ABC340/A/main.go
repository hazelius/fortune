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

	a, b, d := readInt(), readInt(), readInt()
	for i := a; i <= b; i += d {
		fmt.Fprint(out, i, " ")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
