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
	sc.Split(bufio.ScanWords)

	s := readString()
	r := regexp.MustCompile(`[^2]`)
	s = r.ReplaceAllString(s, "")
	fmt.Fprint(out, s)
}

func main() {
	run(os.Stdin, os.Stdout)
}
