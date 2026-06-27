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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}
	t, d, l, r := h-1, 0, w-1, 0
	for i, s := range ss {
		for j, c := range s {
			if c == '.' {
				continue
			}
			if i < t {
				t = i
			}
			if i > d {
				d = i
			}
			if j < l {
				l = j
			}
			if j > r {
				r = j
			}
		}
	}
	for i := t; i <= d; i++ {
		fmt.Fprintln(out, ss[i][l:r+1])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
