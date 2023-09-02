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

type point struct {
	x int
	y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()

	pmap := make(map[point]bool)
	for i := 0; i < n; i++ {
		a, b, c, d := readInt(), readInt(), readInt(), readInt()
		for x := a; x < b; x++ {
			for y := c; y < d; y++ {
				pmap[point{x, y}] = true
			}
		}
	}

	fmt.Fprint(out, len(pmap))
}

func main() {
	run(os.Stdin, os.Stdout)
}
