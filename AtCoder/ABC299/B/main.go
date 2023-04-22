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
	cs := make([]int, n)
	for i := range cs {
		cs[i] = readInt()
	}
	rs := make([]int, n)
	for i := range rs {
		rs[i] = readInt()
	}
	crmap := make(map[int]int)
	for i := 0; i < n; i++ {
		c := cs[i]
		r, ok := crmap[c]
		if !ok {
			crmap[c] = rs[i]
		} else if r < rs[i] {
			crmap[c] = rs[i]
		}
	}

	m, ok := crmap[t]
	if !ok {
		m = crmap[cs[0]]
	}

	for i, r := range rs {
		if r == m {
			fmt.Fprint(out, i+1)
			return
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
