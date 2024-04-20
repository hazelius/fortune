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

	ans := make([][]int, 0, n)

	for idx := 0; idx < n; {
		a := as[idx]
		if a == idx+1 {
			idx++
			continue
		}
		ans = append(ans, []int{idx + 1, a})
		as[idx], as[a-1] = as[a-1], as[idx]
		if as[idx] == idx+1 {
			idx++
		}
	}

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
