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

	n, k, x := readInt(), readInt(), readInt()
	as := make([]int, n+1)
	for i := range as {
		if i == k {
			as[i] = x
			continue
		}
		as[i] = readInt()
	}
	ans := fmt.Sprintf("%v", as)
	fmt.Fprint(out, ans[1:len(ans)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
