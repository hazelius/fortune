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
	as := make([]string, n)
	for i := range as {
		as[i] = readString()
	}
	for i := len(as) - 1; i >= 0; i-- {
		fmt.Fprintln(out, as[i])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
