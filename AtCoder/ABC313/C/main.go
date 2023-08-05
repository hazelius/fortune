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
	sum := 0
	for _, a := range as {
		sum += a
	}

	av := sum / n
	pl := 0
	mn := 0
	for _, a := range as {
		if a < av {
			pl += av - a
		} else if a > av+1 {
			mn += a - (av + 1)
		}
	}
	ans := pl
	if pl < mn {
		ans = mn
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
