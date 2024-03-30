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

	s := readString()
	ans := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			ans[s[i:j]] = true
		}
	}
	fmt.Fprint(out, len(ans))
}

func main() {
	run(os.Stdin, os.Stdout)
}
