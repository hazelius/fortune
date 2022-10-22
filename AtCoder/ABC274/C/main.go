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
	amap[1] = 0
	for i := 0; i < n; i++ {
		a := readInt()
		cnt, ok := amap[a]
		if ok {
			cnt++
			idx := (i + 1) * 2
			amap[idx] = cnt
			amap[idx+1] = cnt
		}
	}

	for i := 1; i <= 2*n+1; i++ {
		a, ok := amap[i]
		if ok {
			fmt.Fprintln(out, a)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
