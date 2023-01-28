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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s, t := readString(), readString()
	statusMap := make(map[int]bool)
	s2 := s[len(s)-len(t):]
	for i := range s2 {
		if s2[i] == '?' || t[i] == '?' {
			continue
		}
		if s2[i] != t[i] {
			statusMap[i] = true
		}
	}

	if len(statusMap) == 0 {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}

	for i := 0; i < len(t); i++ {
		if s[i] == '?' || t[i] == '?' {
			delete(statusMap, i)
		} else if s[i] != t[i] {
			statusMap[i] = true
		} else {
			delete(statusMap, i)
		}
		if len(statusMap) > 0 {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
