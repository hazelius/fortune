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
	as := make([][]int, n)
	for i := range as {
		as[i] = []int{readInt(), readInt()}
	}

	ans := 0
	pre := 0
	for _, a := range as {
		t, v := a[0], a[1]
		ans = max(0, ans-(t-pre))
		ans += v
		pre = t
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
