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

	x, c := readInt(), readInt()
	a := x / (1000 + c)
	fmt.Fprint(out, a*1000)
}

func main() {
	run(os.Stdin, os.Stdout)
}
