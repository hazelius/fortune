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

	n, k := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	ans := make([]int, n)
	cnt := 0
	for i := k; i < len(as); i++ {
		ans[cnt] = as[i]
		cnt++
	}

	s := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, s[1:len(s)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
