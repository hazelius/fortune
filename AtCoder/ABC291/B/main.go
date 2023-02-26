package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	as := make([]int, n*5)
	for i := range as {
		as[i] = readInt()
	}

	sort.Ints(as)
	sum := 0
	for i := n; i < (n*5 - n); i++ {
		sum += as[i]
	}

	fmt.Fprintf(out, "%.15f", float64(sum)/(float64(n)*3))
}

func main() {
	run(os.Stdin, os.Stdout)
}
