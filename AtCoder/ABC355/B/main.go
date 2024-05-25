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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	bs := make([]int, m)
	for i := range bs {
		bs[i] = readInt()
	}

	sort.Ints(as)
	sort.Ints(bs)

	ai, bi := 0, 0
	flg := false
	for ai < n {
		if bi >= m || as[ai] < bs[bi] {
			if flg {
				fmt.Fprint(out, "Yes")
				return
			}
			ai++
			flg = true
		} else {
			bi++
			flg = false
		}
	}

	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
