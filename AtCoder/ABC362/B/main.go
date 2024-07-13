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

	a := vector{readInt(), readInt()}
	b := vector{readInt(), readInt()}
	c := vector{readInt(), readInt()}

	a2 := a.minus(b)
	c2 := c.minus(b)
	cross1 := a2.cross(c2)

	b2 := b.minus(c)
	a2 = a.minus(c)
	cross2 := b2.cross(a2)

	b2 = b.minus(a)
	c2 = c.minus(a)
	cross3 := b2.cross(c2)

	if cross1 == 0 {
		fmt.Fprint(out, "Yes")
	} else if cross2 == 0 {
		fmt.Fprint(out, "Yes")
	} else if cross3 == 0 {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")
	}
}

type vector struct {
	x int
	y int
}

func (v vector) minus(m vector) vector {
	return vector{x: v.x - m.x, y: v.y - m.y}
}

func (v vector) cross(a vector) int {
	return v.x*a.x + v.y*a.y
}

func main() {
	run(os.Stdin, os.Stdout)
}
