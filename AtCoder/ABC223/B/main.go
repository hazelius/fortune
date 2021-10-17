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
	maxans, minans := "", ""
	for i := range s {
		if i == 0 {
			maxans, minans = s, s
			continue
		}

		tmp := s[i:] + s[:i]
		if tmp > maxans {
			maxans = tmp
		}
		if tmp < minans {
			minans = tmp
		}
	}
	return minans + " " + maxans
}

func main() {
	fmt.Println(run(os.Stdin))
}
