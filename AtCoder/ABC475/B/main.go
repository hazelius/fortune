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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	ans1, ans10, ans100 := 0, 0, 0
	for _, a := range as {
		turi := a % 1000
		if turi > 0 {
			turi = 1000 - turi
		}
		ans1 += turi % 10
		ans10 += (turi % 100) / 10
		ans100 += turi / 100
	}

	fmt.Fprint(out, ans1, ans10, ans100)
}

func main() {
	run(os.Stdin, os.Stdout)
}
