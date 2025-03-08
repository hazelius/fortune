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
	bs := make([]int, n)
	for i := range bs {
		bs[i] = readInt()
	}
	ws := make([]int, m)
	for i := range ws {
		ws[i] = readInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(bs)))
	sort.Sort(sort.Reverse(sort.IntSlice(ws)))

	mv := 0
	for i, v := range ws {
		if v > 0 {
			mv += v
		}
		ws[i] = mv
	}

	ans := 0
	bm := 0
	for i, v := range bs {
		bm += v
		wm := ws[len(ws)-1]
		if i < len(ws) {
			wm = ws[i]
		}
		ans = max(ans, bm+wm)
	}

	fmt.Fprint(out, ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
