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

	a := readInt()
	if 400%a > 0 {
		fmt.Fprint(out, -1)
		return
	}
	fmt.Fprint(out, 400/a)
}

func main() {
	run(os.Stdin, os.Stdout)
}
