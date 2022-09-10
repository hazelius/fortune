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
	sc.Split(bufio.ScanWords)

	s, t := readString(), readString()

	if strings.HasPrefix(t, s) {
		fmt.Fprint(out, "Yes")

	} else {
		fmt.Fprint(out, "No")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
