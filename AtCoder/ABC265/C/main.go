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

	h, w := readInt(), readInt()
	gs := make([]string, h)
	for i := range gs {
		gs[i] = readString()
	}

	used := make([][]bool, h)
	for i := range used {
		used[i] = make([]bool, w)
	}

	x, y := 0, 0
	for {
		if used[x][y] {
			fmt.Fprint(out, -1)
			return
		}
		used[x][y] = true

		g := gs[x][y]
		xn, yn := x, y
		switch g {
		case 'U':
			xn--
		case 'D':
			xn++
		case 'L':
			yn--
		case 'R':
			yn++
		}
		if xn < 0 || yn < 0 || xn >= h || yn >= w {
			break
		}
		x, y = xn, yn
	}

	fmt.Fprintf(out, "%d %d", x+1, y+1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
