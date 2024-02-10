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
	as := make([]int, 0, n)
	for i := 0; i < n; i++ {
		switch readInt() {
		case 1:
			as = append(as, readInt())
		case 2:
			idx := len(as) - readInt()
			fmt.Fprintln(out, as[idx])
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
