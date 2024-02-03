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

	min := 0
	ans := 0
	for _, a := range as {
		ans += a
		if ans < min {
			min = ans
		}
	}

	fmt.Fprint(out, ans-min)
}

func main() {
	run(os.Stdin, os.Stdout)
}
