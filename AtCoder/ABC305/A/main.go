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
	a := n % 5
	if a > 2 {
		fmt.Fprint(out, n+(5-a))
	} else {
		fmt.Fprint(out, n-a)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
