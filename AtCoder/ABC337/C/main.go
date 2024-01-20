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
		as[i] = -1
	}

	top := 0
	for i := range as {
		a := readInt() - 1
		if a < 0 {
			top = i
			continue
		}
		as[a] = i
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = top
		top = as[top]
	}

	for _, v := range ans {
		fmt.Fprint(out, v+1, " ")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
