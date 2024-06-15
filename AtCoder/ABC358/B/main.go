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

	n, a := readInt(), readInt()
	s := 0
	for i := 0; i < n; i++ {
		t := readInt()
		if t <= s {
			s += a
		} else {
			s = t + a
		}
		fmt.Fprintln(out, s)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
