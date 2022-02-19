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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	x1, y1, x2, y2 := readInt(), readInt(), readInt(), readInt()
	ps := [][]int{
		{x1 + 2, y1 + 1}, {x1 + 2, y1 - 1},
		{x1 - 2, y1 + 1}, {x1 - 2, y1 - 1},
		{x1 + 1, y1 + 2}, {x1 + 1, y1 - 2},
		{x1 - 1, y1 + 2}, {x1 - 1, y1 - 2},
	}

	for _, p := range ps {
		x := p[0] - x2
		y := p[1] - y2
		if x*x+y*y == 5 {
			return "Yes"
		}
	}
	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
