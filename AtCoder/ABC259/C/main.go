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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s, t := readString(), readString()
	if len(s) > len(t) {
		return "No"
	}

	si := 0
	for ti := 0; ti < len(t); ti++ {
		if si < len(s) && s[si] == t[ti] {
			si++
			continue
		}
		if si < 2 {
			return "No"
		}
		if s[si-2] == s[si-1] && s[si-1] == t[ti] {
			continue
		}
		return "No"
	}

	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
