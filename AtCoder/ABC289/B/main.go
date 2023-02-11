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
	as := make([]int, n+1)
	for i := 0; i < m; i++ {
		a := readInt()
		as[a] = a + 1
	}

	i := 1
	for {
		i = f(i, as, out)
		if i > n {
			break
		}
	}
}

func f(i int, as []int, out io.Writer) int {
	if as[i] == 0 {
		fmt.Fprintf(out, "%v ", i)
		return i + 1
	}
	next := f(as[i], as, out)
	fmt.Fprintf(out, "%v ", i)
	return next
}

func main() {
	run(os.Stdin, os.Stdout)
}
