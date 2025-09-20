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

	n, m, k := readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := 0; i < k; i++ {
		a, _ := readInt()-1, readInt()
		as[a]++
		if as[a] == m {
			fmt.Fprint(out, a+1, " ")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
