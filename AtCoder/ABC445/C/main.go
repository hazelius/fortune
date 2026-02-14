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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt() - 1
	}

	for i := 0; i < 100; i++ {
		tmp := make([]int, n)
		for j := range tmp {
			tmp[j] = as[as[as[as[as[as[as[as[as[as[j]]]]]]]]]]
		}
		as = tmp
	}

	for _, a := range as {
		fmt.Fprint(out, a+1, " ")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
