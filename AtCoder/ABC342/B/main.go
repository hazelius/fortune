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
	mapa := make(map[int]int)
	for i := 0; i < n; i++ {
		p := readInt()
		mapa[p] = i
	}
	q := readInt()
	for i := 0; i < q; i++ {
		a, b := readInt(), readInt()
		if mapa[a] > mapa[b] {
			fmt.Fprintln(out, b)
		} else {
			fmt.Fprintln(out, a)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
