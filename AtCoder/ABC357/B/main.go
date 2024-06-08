package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
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
	countup := 0
	countlow := 0
	for _, v := range s {
		if unicode.IsUpper(v) {
			countup++
		} else {
			countlow++

		}
	}
	if countlow < countup {
		fmt.Fprint(out, strings.ToUpper(s))
		return
	}

	fmt.Fprint(out, strings.ToLower(s))
}

func main() {
	run(os.Stdin, os.Stdout)
}
