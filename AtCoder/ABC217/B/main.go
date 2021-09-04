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

	m := map[string]bool{"ABC": false, "ARC": false, "AGC": false, "AHC": false}
	m[readString()] = true
	m[readString()] = true
	m[readString()] = true
	for k, v := range m {
		if !v {
			return k
		}
	}
	return ""
}

func main() {
	fmt.Println(run(os.Stdin))
}
