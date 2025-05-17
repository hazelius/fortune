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

	_, _, n := readInt(), readInt(), readInt()
	vmap := make(map[int]map[int]bool)
	hmap := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		x, y := readInt(), readInt()
		if _, ok := vmap[x]; !ok {
			vmap[x] = make(map[int]bool)
		}
		if _, ok := hmap[y]; !ok {
			hmap[y] = make(map[int]bool)
		}
		vmap[x][y] = true
		hmap[y][x] = true
	}

	xdelmap := make(map[int]bool)
	ydelmap := make(map[int]bool)
	q := readInt()
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x := readInt()
			ar, ok := vmap[x]
			ydelmap[x] = true
			if ok {
				for a := range ar {
					if xdelmap[a] {
						delete(ar, a)
					}
				}
				delete(vmap, x)
			}
			fmt.Fprintln(out, len(ar))
		case 2:
			y := readInt()
			ar, ok := hmap[y]
			xdelmap[y] = true
			if ok {
				for a := range ar {
					if ydelmap[a] {
						delete(ar, a)
					}
				}
				delete(hmap, y)
			}
			fmt.Fprintln(out, len(ar))
		}

	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
