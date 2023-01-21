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

	n, p, q, r, _ := readInt(), readInt()-1, readInt()-1, readInt()-1, readInt()-1
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	d := r - p
	for i := p; i <= q; i++ {
		as[i], as[i+d] = as[i+d], as[i]
	}
	ans := fmt.Sprintf("%v", as)
	fmt.Fprint(out, ans[1:len(ans)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
