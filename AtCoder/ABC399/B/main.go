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

	n := readInt()
	as := make([][]int, n)
	for i := range as {
		as[i] = []int{i, readInt()}
	}
	sort.Slice(as, func(i, j int) bool {
		return as[i][1] > as[j][1]
	})
	idx := 0
	for i, a := range as {
		if i > 0 && a[1] != as[i-1][1] {
			idx = i
		}
		as[i] = append(as[i], idx)
	}
	ans := make([]int, n)
	for _, a := range as {
		ans[a[0]] = a[2] + 1
	}

	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
