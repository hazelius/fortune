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
	as := make([][]int, n)
	for i := range as {
		as[i] = make([]int, i+1)
		for j := range as[i] {
			as[i][j] = readInt()
		}
	}

	ans := 1
	for i := 1; i <= n; i++ {
		c, d := ans-1, i-1
		if c < d {
			c, d = d, c
		}
		ans = as[c][d]
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
