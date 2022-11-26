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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}
	ts := make([]string, h)
	for i := range ts {
		ts[i] = readString()
	}

	ssw := make([][]rune, w)
	for i := range ssw {
		ssw[i] = make([]rune, h)
	}
	for i, s := range ss {
		for j, c := range s {
			ssw[j][i] = c
		}
	}
	ssstr := make([]string, w)
	for i, s := range ssw {
		ssstr[i] = string(s)
	}

	tsw := make([][]rune, w)
	for i := range tsw {
		tsw[i] = make([]rune, h)
	}
	for i, s := range ts {
		for j, c := range s {
			tsw[j][i] = c
		}
	}
	tsstr := make([]string, w)
	for i, s := range tsw {
		tsstr[i] = string(s)
	}

	sort.Strings(ssstr)
	sort.Strings(tsstr)

	for i := range ssstr {
		if ssstr[i] != tsstr[i] {
			fmt.Fprint(out, "No")
			return
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
