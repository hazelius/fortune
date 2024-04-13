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
	cmap := make(map[rune]int)
	for _, c := range s {
		cmap[c]++
	}
	cntmap := make(map[int]int)
	for _, cnt := range cmap {
		cntmap[cnt]++
	}
	for _, v := range cntmap {
		if v != 2 {
			fmt.Fprint(out, "No")
			return
		}
	}
	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
