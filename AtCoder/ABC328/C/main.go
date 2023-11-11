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

	n, q := readInt(), readInt()
	s := readString()

	ds := make([]int, n)
	for i := range ds {
		if i == 0 {
			continue
		}
		ds[i] = ds[i-1]
		if s[i] == s[i-1] {
			ds[i]++
		}
	}

	for i := 0; i < q; i++ {
		l, r := readInt()-1, readInt()-1

		fmt.Fprintln(out, ds[r]-ds[l])
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
