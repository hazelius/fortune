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

	n, t, a := readInt(), readInt(), readInt()
	if t > n/2 || a > n/2 {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")

	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
