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
	as := make([]int, n)
	idx := -1
	for i := 0; i < n*7; i++ {
		if i%7 == 0 {
			idx++
		}
		as[idx] += readInt()
	}
	ans := fmt.Sprintf("%v", as)
	fmt.Fprint(out, ans[1:len(ans)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
