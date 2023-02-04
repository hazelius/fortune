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
	for i := 0; i < n; i++ {
		a, b := readInt(), readInt()
		fmt.Fprintln(out, a+b)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
