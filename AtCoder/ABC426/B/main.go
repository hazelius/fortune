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
	smap := make(map[rune]int)
	for _, v := range s {
		smap[v]++
	}
	for k, v := range smap {
		if v == 1 {
			fmt.Fprint(out, string(k))
			return
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
