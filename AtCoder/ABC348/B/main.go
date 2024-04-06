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
	xys := make([][]int, n)
	for i := range xys {
		xys[i] = []int{readInt(), readInt()}
	}
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, n)
	}

	for a, d := range dis {
		mx := 0
		ans := 0
		for b, v := range d {
			if a == b {
				continue
			}

			if v == 0 {
				lx := (xys[a][0] - xys[b][0])
				ly := (xys[a][1] - xys[b][1])
				v = lx*lx + ly*ly
				dis[a][b] = v
				dis[b][a] = v
			}
			if mx < v {
				mx = v
				ans = b
			}
		}
		fmt.Fprintln(out, ans+1)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
