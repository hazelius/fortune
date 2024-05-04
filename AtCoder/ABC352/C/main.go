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
	ans := 0
	m := 0
	for i := 0; i < n; i++ {
		a, b := readInt(), readInt()
		ans += a
		sa := b - a
		if sa > m {
			m = sa
		}
	}
	fmt.Fprint(out, ans+m)
}

func main() {
	run(os.Stdin, os.Stdout)
}
