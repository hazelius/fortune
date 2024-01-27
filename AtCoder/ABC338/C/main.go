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
	qs := make([]int, n)
	for i := range qs {
		qs[i] = readInt()
	}
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	bs := make([]int, n)
	for i := range bs {
		bs[i] = readInt()
	}

	ans := 0
	flg := true
	for i := 0; flg; i++ {
		bmax := -1
		for i, q := range qs {
			if bs[i] == 0 {
				continue
			}
			btmp := q / bs[i]
			if bmax < 0 {
				bmax = btmp
			} else {
				bmax = min(bmax, btmp)
			}
		}

		ans = max(ans, i+bmax)

		for i := range qs {
			qs[i] -= as[i]
			if qs[i] < 0 {
				flg = false
				break
			}
		}
	}

	fmt.Fprint(out, ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
