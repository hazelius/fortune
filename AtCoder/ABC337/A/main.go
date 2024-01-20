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
	x, y := 0, 0
	for i := 0; i < n; i++ {
		x += readInt()
		y += readInt()
	}
	if x == y {
		fmt.Fprint(out, "Draw")
	} else if x > y {
		fmt.Fprint(out, "Takahashi")
	} else {
		fmt.Fprint(out, "Aoki")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
