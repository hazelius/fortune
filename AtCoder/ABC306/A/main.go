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
	s := readString()
	as := make([]rune, 2*n)

	i := 0
	for _, c := range s {
		as[i] = c
		i++
		as[i] = c
		i++
	}
	fmt.Fprint(out, string(as))
}

func main() {
	run(os.Stdin, os.Stdout)
}
