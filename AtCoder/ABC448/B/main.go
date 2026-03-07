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
	cs := make([]int, m)
	for i := range cs {
		cs[i] = readInt()
	}
	ans := 0
	for i := 0; i < n; i++ {
		a, b := readInt()-1, readInt()
		if b < cs[a] {
			ans += b
			cs[a] -= b
		} else {
			ans += cs[a]
			cs[a] = 0
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
