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
	for i := range as {
		as[i] = readInt()
	}
	bs := make([]int, n-1)
	for i := range bs {
		bs[i] = as[i] * as[i+1]
	}
	ans := fmt.Sprintf("%v", bs)
	fmt.Fprint(out, ans[1:len(ans)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
