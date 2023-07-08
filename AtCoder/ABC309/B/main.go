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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]string, n)
	for i := range as {
		as[i] = readString()
	}

	var r string
	for i, a := range as {
		if i == 0 {
			r = a[n-1:]
			as[i] = as[i+1][:1] + a[:n-1]
		} else if i == n-1 {
			as[i] = a[1:] + r
		} else {
			t := a[n-1:]
			as[i] = as[i+1][:1] + a[1:n-1] + r
			r = t
		}
	}

	for _, a := range as {
		fmt.Fprintln(out, a)
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
