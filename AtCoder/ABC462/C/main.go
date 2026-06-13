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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readInt()
	xys := make([][]int, n)
	for i := range xys {
		xys[i] = []int{readInt(), readInt()}
	}
	sort.Slice(xys, func(x, y int) bool {
		return xys[x][0] < xys[y][0]
	})

	ans := 0
	miny := 999999
	for _, xy := range xys {
		y := xy[1]
		if miny > y {
			ans++
			miny = y
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
