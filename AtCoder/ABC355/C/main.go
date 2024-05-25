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

	n, t := readInt(), readInt()

	yoko := make([]map[int]bool, n)
	for i := range yoko {
		yoko[i] = make(map[int]bool)
	}
	tate := make([]map[int]bool, n)
	for i := range tate {
		tate[i] = make(map[int]bool)
	}
	naname := make([]map[int]bool, 2)
	for i := range naname {
		naname[i] = make(map[int]bool)
	}
	for i := 0; i < t; i++ {
		a := readInt() - 1
		x := a / n
		y := (a % n)
		yoko[x][y] = true
		tate[y][x] = true
		if x == y {
			naname[0][x] = true
		}
		if x+y == n-1 {
			naname[1][x] = true
		}
		if len(yoko[x]) == n || len(tate[y]) == n || len(naname[0]) == n || len(naname[1]) == n {
			fmt.Fprint(out, i+1)
			return
		}
	}
	fmt.Fprint(out, -1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
