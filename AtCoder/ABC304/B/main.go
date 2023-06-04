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

	for i := 1000; i <= 1000000000; i *= 10 {
		if n <= i {
			break
		}
		n -= n % (i / 100)
	}
	fmt.Fprint(out, n)
}

func main() {
	run(os.Stdin, os.Stdout)
}
