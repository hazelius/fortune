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

	n, k := readInt(), readInt()
	acnt := make([]int, k)
	for i := 0; i < n; i++ {
		a := readInt() - 1
		acnt[a]++
	}
	amap := make(map[int]int)
	for _, cnt := range acnt {
		amap[cnt]++
	}
	maxScnt := 0
	ansCcnt := 0
	for scnt, ccnt := range amap {
		if maxScnt < scnt {
			ansCcnt = ccnt
			maxScnt = scnt
		}
	}
	if v, ok := amap[maxScnt-1]; ok {
		ansCcnt += v
	}
	fmt.Fprint(out, ansCcnt)
}

func main() {
	run(os.Stdin, os.Stdout)
}
