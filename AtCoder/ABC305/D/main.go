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

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	stime := make([]int, n)
	for i := 1; i < n; i++ {
		stime[i] = stime[i-1]
		if i%2 == 0 {
			stime[i] += as[i] - as[i-1]
		}
	}

	q := readInt()
	for i := 0; i < q; i++ {
		l, r := readInt(), readInt()
		lp := sort.Search(n, func(j int) bool { return as[j] > l }) - 1
		rp := sort.Search(n, func(j int) bool { return as[j] > r }) - 1
		ans := stime[rp] - stime[lp]
		if lp%2 > 0 {
			ans -= l - as[lp]
		}
		if rp%2 > 0 {
			ans += r - as[rp]
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
