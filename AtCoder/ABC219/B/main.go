package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	ss := make([]string, 3)
	for i := range ss {
		ss[i] = readString()
	}
	t := readString()

	ans := make([]string, len(t))
	for i, c := range t {
		num := int(c - '1')
		ans[i] = ss[num]
	}
	return strings.Join(ans, "")
}

func main() {
	fmt.Println(run(os.Stdin))
}
