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

	n := readInt()
	amap := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a] = true
	}
	m := readInt()
	bmap := make(map[int]bool)
	for i := 0; i < m; i++ {
		b := readInt()
		bmap[b] = true
	}
	l := readInt()
	cmap := make(map[int]bool)
	for i := 0; i < l; i++ {
		c := readInt()
		cmap[c] = true
	}

	ansmap := make(map[int]bool)
	for a := range amap {
		for b := range bmap {
			for c := range cmap {
				abc := a + b + c
				ansmap[abc] = true
			}
		}

	}

	q := readInt()
	for i := 0; i < q; i++ {
		x := readInt()
		if ansmap[x] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
