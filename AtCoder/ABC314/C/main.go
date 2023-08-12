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

	n, m := readInt(), readInt()
	s := []byte(readString())
	tmpc := make([]byte, m)
	firstIdx := make([]int, m)

	for i := 0; i < n; i++ {
		cha := s[i]
		c := readInt() - 1
		if tmpc[c] == 0 {
			firstIdx[c] = i
		} else {
			s[i] = tmpc[c]
		}
		tmpc[c] = cha
	}

	for c, idx := range firstIdx {
		s[idx] = tmpc[c]
	}
	fmt.Fprint(out, string(s))
}

func main() {
	run(os.Stdin, os.Stdout)
}
