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

	x, k := readInt(), readInt()

	for i := 0; i < k; i++ {
		a := x % 10
		x /= 10
		if a >= 5 {
			x++
		}
	}

	for i := 0; i < k; i++ {
		x *= 10
	}

	fmt.Fprint(out, x)
}

func main() {
	run(os.Stdin, os.Stdout)
}
