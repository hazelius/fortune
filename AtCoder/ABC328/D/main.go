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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()

	l := len(s)
	s = strings.Replace(s, "ABC", "", -1)
	for l != len(s) {
		l = len(s)
		s = strings.Replace(s, "ABC", "", -1)
	}
	fmt.Fprint(out, s)

}

func main() {
	run(os.Stdin, os.Stdout)
}
