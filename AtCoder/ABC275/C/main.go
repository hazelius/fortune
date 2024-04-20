// https://atcoder.jp/contests/abc275/tasks/abc275_c

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

type vector struct {
	x int
	y int
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	maps := make(map[vector]bool)
	for i := 0; i < 9; i++ {
		s := readString()
		for j, c := range s {
			if c == '#' {
				maps[vector{i, j}] = true
			}
		}
	}

	ans := 0
	for p1 := range maps {
		for p2 := range maps {
			if p1.x == p2.x && p1.y == p2.y {
				continue
			}
			xd := p2.x - p1.x
			yd := p2.y - p1.y
			p3 := vector{p2.x - yd, p2.y + xd}
			p4 := vector{p3.x - xd, p3.y - yd}
			if maps[p3] && maps[p4] {
				ans++
			}
		}
	}

	fmt.Fprint(out, ans/4)
}

func main() {
	run(os.Stdin, os.Stdout)
}
