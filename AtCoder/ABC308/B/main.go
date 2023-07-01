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

	n, m := readInt(), readInt()
	cs := make([]string, n)
	for i := range cs {
		cs[i] = readString()
	}
	ds := make([]string, m)
	for i := range ds {
		ds[i] = readString()
	}
	p0 := readInt()
	ps := make([]int, m)
	for i := range ps {
		ps[i] = readInt()
	}
	cmap := make(map[string]int)
	for i, d := range ds {
		cmap[d] = ps[i]
	}

	ans := 0
	for _, c := range cs {
		v, ok := cmap[c]
		if ok {
			ans += v
		} else {
			ans += p0
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
