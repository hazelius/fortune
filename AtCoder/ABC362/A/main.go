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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	r, g, b := readInt(), readInt(), readInt()
	c := readString()
	switch c {
	case "Blue":
		fmt.Fprint(out, min(r, g))
	case "Green":
		fmt.Fprint(out, min(r, b))
	default:
		fmt.Fprint(out, min(g, b))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
