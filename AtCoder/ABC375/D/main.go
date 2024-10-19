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
	cmap := make(map[rune][]int)
	for i, c := range s {
		cs, ok := cmap[c]
		if !ok {
			cs = make([]int, 0)
		}
		cs = append(cs, i)
		cmap[c] = cs
	}

	ans := 0
	for _, cs := range cmap {
		pre := 0
		for i, c := range cs {
			if i == 0 {
				pre = c
				continue
			}
			ans += (c - pre - 1) * len(cs[i:])
			ans += (c - pre) * len(cs[:i-1]) * len(cs[i:])
			pre = c
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
