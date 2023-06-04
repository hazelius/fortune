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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	ss := make([]string, n)
	startIdx := 0
	min := 1000000000
	for i := range as {
		ss[i] = readString()
		as[i] = readInt()
		if min > as[i] {
			min = as[i]
			startIdx = i
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, ss[(startIdx)%n])
		startIdx++
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
