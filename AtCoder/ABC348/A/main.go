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
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Fprint(out, "x")
		} else {
			fmt.Fprint(out, "o")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
