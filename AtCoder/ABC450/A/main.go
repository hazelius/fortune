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
	for i := n; i > 0; i-- {
		fmt.Fprint(out, i)
		if i != 1 {
			fmt.Fprint(out, ",")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
