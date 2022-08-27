// C - Convex Quadrilateral
// https://atcoder.jp/contests/abc266/tasks/abc266_c
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

type vector struct {
	x int
	y int
}

func (v vector) minus(m vector) vector {
	return vector{x: v.x - m.x, y: v.y - m.y}
}
func (v vector) cross(a vector) int {
	return v.x*a.y - v.y*a.x
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	vs := make([]vector, 4)
	for i := 0; i < 4; i++ {
		vs[i] = vector{x: readInt(), y: readInt()}
	}

	for i := 0; i < 4; i++ {
		a := vs[i]
		b := vs[(i+1)%4]
		c := vs[(i+2)%4]

		a2 := a.minus(b)
		c2 := c.minus(b)
		cross := a2.cross(c2)
		if cross > 0 {
			fmt.Fprint(out, "No")
			return
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
