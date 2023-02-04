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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}
	ss2 := ss[0:k]
	sort.Strings(ss2)
	for _, v := range ss2 {
		fmt.Fprintln(out, v)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
