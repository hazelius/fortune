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

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		s := readString()
		bs := make([]byte, n)
		idx := -1
		flg := false
		for j := range s {
			if flg {
				bs[j] = s[j]
				continue
			}
			if idx < 0 {
				if j < n-1 && s[j] > s[j+1] {
					idx = j
					bs[j] = s[j+1]
				} else {
					bs[j] = s[j]
				}
			} else {
				if j == n-1 || s[j+1] > s[idx] {
					flg = true
					bs[j] = s[idx]
				} else {
					bs[j] = s[j+1]
				}
			}
		}
		fmt.Fprintln(out, string(bs))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
