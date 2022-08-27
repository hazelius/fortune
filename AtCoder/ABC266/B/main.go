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

	d := 998244353
	n := readInt()
	x := n % d
	if x < 0 {
		x = d + x
	}

	fmt.Fprint(out, x)
}

func main() {
	run(os.Stdin, os.Stdout)
}
