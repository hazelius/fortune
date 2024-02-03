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
	ss := strings.Split(s, ".")
	fmt.Fprint(out, ss[len(ss)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
