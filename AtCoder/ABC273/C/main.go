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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	ans := make([]int, n)
	sort.Sort(sort.Reverse(sort.IntSlice(as)))

	cnt := 0
	for i, a := range as {
		if i == 0 {
			ans[cnt] = 1
			continue
		}
		if a == as[i-1] {
			ans[cnt]++
			continue
		}
		cnt++
		ans[cnt]++
	}

	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
