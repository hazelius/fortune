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
	ls := make([][]int, n)
	for i := range ls {
		l := readInt()
		ls[i] = make([]int, l)
		for j := range ls[i] {
			ls[i][j] = readInt()
		}
	}
	for i := 0; i < q; i++ {
		s, t := readInt(), readInt()
		fmt.Fprintln(out, ls[s-1][t-1])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
