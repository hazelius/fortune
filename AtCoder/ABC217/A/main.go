package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	s, t := readString(), readString()
	ar := []string{s, t}
	sort.Strings(ar)
	if ar[0] == s {
		return "Yes"
	}
	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
