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

type vec struct {
	x int
	y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	abmap := make(map[int]map[int]vec)
	for i := 0; i < m; i++ {
		a, b, x, y := readInt()-1, readInt()-1, readInt(), readInt()
		if _, ok := abmap[a]; !ok {
			abmap[a] = make(map[int]vec)
		}
		abmap[a][b] = vec{x, y}
		if _, ok := abmap[b]; !ok {
			abmap[b] = make(map[int]vec)
		}
		abmap[b][a] = vec{-x, -y}
	}

	ansmap := make(map[int]vec)
	ansmap[0] = vec{0, 0}
	f(0, abmap, ansmap)

	for i := 0; i < n; i++ {
		xy, ok := ansmap[i]
		if !ok {
			fmt.Fprintln(out, "undecidable")
		} else {
			fmt.Fprintln(out, xy.x, xy.y)
		}
	}

}

func f(i int, abmap map[int]map[int]vec, ansmap map[int]vec) {
	dirmap, ok := abmap[i]
	if !ok {
		return
	}

	gen := ansmap[i]

	for b, dir := range dirmap {
		if _, ok := ansmap[b]; ok {
			continue
		}
		ansmap[b] = vec{gen.x + dir.x, gen.y + dir.y}
		f(b, abmap, ansmap)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
