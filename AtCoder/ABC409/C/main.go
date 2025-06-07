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

	n, l := readInt(), readInt()
	ds := make([]int, n-1)
	for i := range ds {
		ds[i] = readInt()
	}
	if l%3 > 0 {
		fmt.Fprint(out, 0)
		return
	}

	idx := 0
	ps := make([]int, l)
	ps[idx]++
	for _, d := range ds {
		idx = (idx + d) % l
		ps[idx]++
	}
	ans := 0
	cnt := l / 3
	for i := 0; i < cnt; i++ {
		ans += ps[i] * ps[i+cnt] * ps[i+2*cnt]
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
