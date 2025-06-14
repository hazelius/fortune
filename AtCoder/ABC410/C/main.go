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

	n, q := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = i + 1
	}

	b := 0
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			p, x := readInt()-1, readInt()
			p = (p + b) % n
			as[p] = x
		case 2:
			p := readInt() - 1
			p = (p + b) % n
			fmt.Fprintln(out, as[p])
		case 3:
			b += readInt()
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
