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
		m -= readInt()
		if m == 0 {
			fmt.Fprint(out, i+1)
			return
		}
		if m < 0 {
			fmt.Fprint(out, i)
			return
		}
	}
	fmt.Fprint(out, n)
}

func main() {
	run(os.Stdin, os.Stdout)
}
