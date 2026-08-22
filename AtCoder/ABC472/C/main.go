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

	n, m, k := readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	sum := 0
	for i, a := range as {
		if sum+a <= k {
			sum += a
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
			as[i] = 0
		}
		if i+1 >= m {
			sum -= as[i+1-m]
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
