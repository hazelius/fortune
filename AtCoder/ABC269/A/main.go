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

	a, b, c, d := readInt(), readInt(), readInt(), readInt()
	fmt.Fprintln(out, (a+b)*(c-d))
	fmt.Fprint(out, "Takahashi")
}

func main() {
	run(os.Stdin, os.Stdout)
}
