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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	amap := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[i] = a
	}

	q := readInt()
	base := 0
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			base = readInt()
			amap = make(map[int]int)
		case 2:
			a, x := readInt(), readInt()
			amap[a-1] += x
		case 3:
			a := readInt()
			fmt.Fprintln(out, base+amap[a-1])
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
