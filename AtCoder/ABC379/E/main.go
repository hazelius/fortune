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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	s := readString()
	ans := make([]int, n+50)
	val := 0
	for i := range s {
		dig, _ := strconv.Atoi(s[i : i+1])
		val += dig * (i + 1)
		idx := len(s) - i - 1
		ans[idx] += val
	}

	amari := 0
	for i, v := range ans {
		v += amari
		ans[i] = v % 10
		amari = v / 10
	}
	flg := false
	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] != 0 {
			flg = true
		}
		if flg {
			fmt.Fprint(out, ans[i])
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
