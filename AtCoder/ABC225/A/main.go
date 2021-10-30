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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	s := readString()
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	if len(m) == 3 {
		return 6
	} else if len(m) == 2 {
		return 3
	}
	return 1
}

func main() {
	fmt.Println(run(os.Stdin))
}
