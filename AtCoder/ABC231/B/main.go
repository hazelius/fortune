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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	smap := make(map[string]int)
	for i := 0; i < n; i++ {
		s := readString()
		smap[s]++
	}
	ans := ""
	m := 0
	for s, p := range smap {
		if p > m {
			ans = s
			m = p
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
