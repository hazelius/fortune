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

	n, x, y := readInt(), readInt(), readInt()
	red := 1
	blue := 0
	for i := n; i > 1; i-- {
		blue += red * x
		red += blue
		blue = blue * y
	}

	return blue
}

func main() {
	fmt.Println(run(os.Stdin))
}
