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

	a, b, c, d, e, f := readInt(), readInt(), readInt(), readInt(), readInt(), readInt()
	g, h, i, j, k, l := readInt(), readInt(), readInt(), readInt(), readInt(), readInt()

	if fb(a, d, g, j) && fb(b, e, h, k) && fb(c, f, i, l) {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")
	}

}

func fb(a, b, c, d int) bool {
	if a > b {
		a, b = b, a
	}
	if c > d {
		c, d = d, c
	}

	if b <= c || d <= a {
		return false
	}
	return true
}

func main() {
	run(os.Stdin, os.Stdout)
}
