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
	cs := make([]int, 9)
	minc := -1
	mincIdx := 0
	for i := range cs {
		c := readInt()
		cs[i] = c
		if minc < 0 || minc >= c {
			minc = c
			mincIdx = i
		}
	}
	cnt := n / minc
	yosan := n % minc
	ans := make([]int, cnt)
	flg := false
	for idx := 0; idx < cnt; idx++ {
		if flg {
			ans[idx] = mincIdx + 1
			continue
		}

		for i := 8; i >= mincIdx; i-- {
			if cs[i] <= minc+yosan {
				ans[idx] = i + 1
				yosan -= cs[i] - minc
				if yosan == 0 || i == mincIdx {
					flg = true
				}
				break
			}
		}
	}

	for _, v := range ans {
		fmt.Fprint(out, v)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
