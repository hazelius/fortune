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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	sumas := make([]int, n+1)
	for i := 0; i < n; i++ {
		sumas[i+1] = sumas[i] + as[i]
	}

	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x := readInt() - 1
			sumas[x+1] += as[x+1] - as[x]
			as[x], as[x+1] = as[x+1], as[x]
		case 2:
			l, r := readInt(), readInt()
			fmt.Fprintln(out, sumas[r]-sumas[l-1])
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
