package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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
	r := regexp.MustCompile(`(W+)(A)`)
	ans := r.ReplaceAllStringFunc(s, func(match string) string {
		w := match[0:strings.Index(match, "A")]
		return "A" + strings.Repeat("C", len(w))
	})

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
