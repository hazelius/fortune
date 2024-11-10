package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	xas := make([][]int, m)
	for i := range xas {
		xas[i] = []int{readInt(), 0}
	}
	for i := range xas {
		xas[i][1] = readInt()
	}
	sort.Slice(xas, func(i, j int) bool {
		return xas[i][0] < xas[j][0]
	})
	if xas[0][0] != 1 {
		fmt.Fprint(out, -1)
		return
	}

	ans := 0
	p := 0
	for i, xa := range xas {
		next := n + 1
		if i < m-1 {
			next = xas[i+1][0]
		}
		p += xa[1]
		d := next - xa[0]
		if d > p {
			fmt.Fprint(out, -1)
			return
		}

		ans += d * (d - 1) / 2
		p -= d
		ans += p * d
	}
	if p > 0 {
		fmt.Fprint(out, -1)
		return
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
