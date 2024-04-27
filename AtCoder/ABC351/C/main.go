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

	idx := 0
	ans := make([]int, n)
	for _, a := range as {
		ans[idx] = a
		idx++
		for idx > 1 && ans[idx-2] == ans[idx-1] {
			idx--
			ans[idx-1] += 1
		}
	}

	fmt.Fprint(out, idx)
}

func main() {
	run(os.Stdin, os.Stdout)
}
