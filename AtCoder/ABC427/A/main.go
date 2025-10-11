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
	sc.Split(bufio.ScanWords)

	s := readString()
	ln := len(s)
	fmt.Fprint(out, s[:ln/2], s[(ln/2)+1:])
}

func main() {
	run(os.Stdin, os.Stdout)
}
