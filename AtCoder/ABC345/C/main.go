package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()
	smap := make(map[rune]int)
	for _, c := range s {
		smap[c]++
	}

	l := len(s)
	ans := 0
	flg := false
	for _, cnt := range smap {
		l -= cnt
		ans += cnt * l
		if cnt > 1 {
			flg = true
		}
	}
	if flg {
		ans++
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
