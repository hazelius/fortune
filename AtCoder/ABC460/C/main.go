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

	n, m := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	bs := make([]int, m)
	for i := range bs {
		bs[i] = readInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(as)))
	sort.Sort(sort.Reverse(sort.IntSlice(bs)))

	ans := 0
	for i, j := 0, 0; i < n && j < m; {
		if as[i]*2 >= bs[j] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
