package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

	s := readString()
	ans := 0

	for idx := 0; idx < len(s)-1; {
		a1, _ := strconv.Atoi(s[idx : idx+1])
		a2, _ := strconv.Atoi(s[idx+1 : idx+2])
		if a1+1 == a2 {
			ans++
			for i, j := idx-1, idx+2; i >= 0 && j < len(s); i, j = i-1, j+1 {
				if s[i] == s[idx] && s[j] == s[idx+1] {
					ans++
				} else {
					idx = j - 2
					break
				}
			}
		}
		idx++
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
