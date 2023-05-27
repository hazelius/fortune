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
	as := make([][]int, m)
	for i := range as {
		as[i] = make([]int, n)
		for j := range as[i] {
			as[i][j] = readInt() - 1
		}
	}

	pairs := make([][]bool, n)
	for i := range pairs {
		pairs[i] = make([]bool, n)
	}
	cnt := 0
	for _, a := range as {
		for i := 0; i < n-1; i++ {
			a1, a2 := a[i], a[i+1]
			if a1 > a2 {
				a1, a2 = a2, a1
			}
			if !pairs[a1][a2] {
				cnt++
				pairs[a1][a2] = true
			}
		}
	}

	fmt.Fprint(out, (n*(n-1)/2)-cnt)
}

func main() {
	run(os.Stdin, os.Stdout)
}
