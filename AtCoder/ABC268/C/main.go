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
	ps := make([]int, n)
	for i := range ps {
		ps[i] = readInt()
	}

	hits := make([]int, n)
	for i, v := range ps {
		m := v - i + 1
		for j := 0; j < 3; j++ {
			m--
			if m < 0 {
				m = n + m
			}
			hits[m]++
		}
	}

	ans := 0
	for _, v := range hits {
		if ans < v {
			ans = v
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
