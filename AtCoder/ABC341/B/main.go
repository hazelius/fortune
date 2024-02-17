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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	for i := 0; i < n-1; i++ {
		s, t := readInt(), readInt()
		as[i+1] += (as[i] / s) * t
	}

	fmt.Fprint(out, as[n-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
