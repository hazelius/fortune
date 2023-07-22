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
		as[i] = readInt() - 1
	}

	ans := make([]int, 0)
	used := make([]bool, n)

	p := 0
	for i := 0; i < n; i++ {
		if used[p] {
			break
		}
		used[p] = true
		ans = append(ans, p)
		p = as[p]
	}

	for i, v := range ans {
		if v == p {
			ans = ans[i:]
			break
		}
	}

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintf(out, "%v ", v+1)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
