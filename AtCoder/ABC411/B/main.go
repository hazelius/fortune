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
	ds := make([]int, n-1)
	for i := range ds {
		ds[i] = readInt()
	}

	for i, d := range ds {
		ans := d
		for j := i + 1; j < n-1; j++ {
			fmt.Fprint(out, ans, " ")
			ans += ds[j]
		}
		fmt.Fprint(out, ans, "\n")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
