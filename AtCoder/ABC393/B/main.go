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
	sc.Split(bufio.ScanWords)

	s := readString()
	as := make([]int, 0)
	bs := make([]int, 0)
	for i, c := range s {
		if c == 'A' {
			as = append(as, i)
		}
		if c == 'B' {
			bs = append(bs, i)
		}
	}

	ans := 0
	for _, a := range as {
		for _, b := range bs {
			if b < a {
				continue
			}
			c := b + (b - a)
			if c < len(s) && s[c] == 'C' {
				ans++
			}
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
