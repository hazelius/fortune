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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	mapcnts := make(map[int]int)
	for _, a := range as {
		cnt, ok := mapcnts[a]
		if ok {
			next, ok := mapcnts[a+1]
			if !ok || next < cnt+1 {
				mapcnts[a+1] = cnt + 1
			}
			delete(mapcnts, a)
		} else {
			if _, ok := mapcnts[a+1]; !ok {
				mapcnts[a+1] = 1
			}
		}
	}

	ans := 0
	for _, cnt := range mapcnts {
		if ans < cnt {
			ans = cnt
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
