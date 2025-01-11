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

	n, r := readInt(), readInt()
	for i := 0; i < n; i++ {
		d, a := readInt(), readInt()
		l, h := 1600, 2799
		if d == 2 {
			l, h = 1200, 2399
		}
		if l <= r && r <= h {
			r += a
		}
	}
	fmt.Fprint(out, r)
}

func main() {
	run(os.Stdin, os.Stdout)
}
