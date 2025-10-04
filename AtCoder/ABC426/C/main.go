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
		as[i] = 1
	}
	min := 0
	for i := 0; i < q; i++ {
		x, y := readInt()-1, readInt()-1
		if x < min {
			fmt.Fprintln(out, 0)
			continue
		}
		cnt := 0
		for ; min <= x; min++ {
			cnt += as[min]
			as[y] += as[min]
		}
		fmt.Fprintln(out, cnt)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
