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
	ps := make([]int, n)
	for i := range ps {
		ps[i] = readInt() - 1
	}

	qs := make([]int, n)
	qs2 := make([]int, n)
	for i := range qs {
		q := readInt() - 1
		qs[q] = i
		qs2[i] = q
	}

	ans := make([]int, n)
	for i := range ans {
		h := qs[i]
		h2 := ps[h]
		s := qs2[h2]
		ans[i] = s + 1
	}

	str := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, str[1:len(str)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
