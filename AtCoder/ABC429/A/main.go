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

	n, m := readInt(), readInt()
	for i := 0; i < n; i++ {
		if i < m {
			fmt.Fprintln(out, "OK")
		} else {
			fmt.Fprintln(out, "Too Many Requests")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
