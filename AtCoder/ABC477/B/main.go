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

	n, d := readInt(), readInt()
	xs := make([]int, n)
	xmap := make(map[int]int)
	for i := range xs {
		x := readInt()
		xs[i] = x
		xmap[x] = i + 1
	}

	sort.Ints(xs)

	ans := make([]int, 0, n)
	for i, x := range xs {
		flg := true
		if i > 0 {
			if x-xs[i-1] < d {
				flg = false
			}
		}
		if i < n-1 {
			if xs[i+1]-x < d {
				flg = false
			}
		}
		if flg {
			ans = append(ans, x)
		}
	}

	fmt.Fprintln(out, len(ans))
	ansarr := make([]int, len(ans))
	for i, x := range ans {
		ansarr[i] = xmap[x]
	}
	sort.Ints(ansarr)
	ansstr := fmt.Sprintf("%v", ansarr)
	fmt.Fprint(out, ansstr[1:len(ansstr)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
