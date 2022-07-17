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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	s := readString()
	smap := make(map[rune]int)
	for _, c := range s {
		smap[c]++
	}
	for k, v := range smap {
		if v == 1 {
			return string(k)
		}
	}
	return "-1"
}

func main() {
	fmt.Println(run(os.Stdin))
}
