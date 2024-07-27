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

	n, x, y := readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	bs := make([]int, n)
	for i := range bs {
		bs[i] = readInt()
	}

	abs := make([][]int, n)
	for i := range abs {
		abs[i] = []int{as[i], bs[i]}
	}

	sort.Slice(abs, func(i, j int) bool {
		return abs[i][0] > abs[j][0]
	})

	asum := 0
	ans := 0
	for i, ab := range abs {
		asum += ab[0]
		ans = i + 1
		if asum > x {
			break
		}
	}

	sort.Slice(abs, func(i, j int) bool {
		return abs[i][1] > abs[j][1]
	})

	asum = 0
	for i, ab := range abs {
		asum += ab[1]
		if asum > y {
			if ans > i+1 {
				ans = i + 1
			}
			break
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
