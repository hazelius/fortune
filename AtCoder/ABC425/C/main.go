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
		as[i] = readInt()
	}

	asum := make([]int, n*2)
	for i := range asum {
		asum[i] = as[i%n]
		if i > 0 {
			asum[i] += asum[i-1]
		}
	}

	idx := 0
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			c := readInt()
			idx = (idx + c) % n
		case 2:
			l, r := readInt()-1, readInt()-1
			if l+idx == 0 {
				fmt.Fprintln(out, asum[r+idx])
			} else {
				fmt.Fprintln(out, asum[r+idx]-asum[l+idx-1])
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
