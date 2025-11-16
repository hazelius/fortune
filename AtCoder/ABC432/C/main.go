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

	n, x, y := readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	sort.Ints(as)
	m := as[0] * y
	if m < as[n-1]*x {
		fmt.Fprint(out, -1)
		return
	}

	sa := y - x
	ans := 0
	for _, a := range as {
		t := m - a*x
		if t%sa != 0 {
			fmt.Fprint(out, -1)
			return
		}

		ans += t / sa
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
