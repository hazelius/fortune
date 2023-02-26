package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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
	r := regexp.MustCompile(`[A-Z]`)
	ans := r.FindStringIndex(s)

	fmt.Fprint(out, ans[0]+1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
