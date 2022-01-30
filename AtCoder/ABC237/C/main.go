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

	s := readString()
	fd, bk := 0, 0
	for ; fd < len(s); fd++ {
		if s[fd] != 'a' {
			break
		}
	}
	if fd == len(s) {
		return "Yes"
	}

	for bk = len(s) - 1; bk > 0; bk-- {
		if s[bk] != 'a' {
			break
		}
	}

	if fd >= len(s)-bk {
		return "No"
	}

	for fd < bk {
		if s[fd] != s[bk] {
			return "No"
		}
		fd++
		bk--
	}

	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
