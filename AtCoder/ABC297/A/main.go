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

	n, d := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	for i := 1; i < n; i++ {
		if as[i]-as[i-1] <= d {
			fmt.Fprint(out, as[i])
			return
		}
	}

	fmt.Fprint(out, -1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
