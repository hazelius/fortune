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
	as := make([]int, m)
	for i := range as {
		as[i] = readInt()
	}

	idx := 0
	for i := 0; i < n; i++ {
		a := as[idx]
		if i >= a {
			idx++
			a = as[idx]
		}
		fmt.Fprintln(out, a-i-1)
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
