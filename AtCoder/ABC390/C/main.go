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

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	hh, hl, wh, wl := 0, h, 0, w
	for i, s := range ss {
		for j, c := range s {
			if c != '#' {
				continue
			}
			if i < hl {
				hl = i
			}
			if i > hh {
				hh = i
			}
			if j < wl {
				wl = j
			}
			if j > wh {
				wh = j
			}
		}
	}

	for i := hl; i <= hh; i++ {
		for j := wl; j <= wh; j++ {
			if ss[i][j] == '.' {
				fmt.Fprint(out, "No")
				return
			}
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
