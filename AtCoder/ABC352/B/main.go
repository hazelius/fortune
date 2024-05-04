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
	ans := make([]int, 0)
	idx := 0
	for i := range t {
		if t[i] == s[idx] {
			idx++
			ans = append(ans, i+1)
		}
	}
	st := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, st[1:len(st)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
