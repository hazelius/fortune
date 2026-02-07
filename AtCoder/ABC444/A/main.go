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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readString()
	if n[0] == n[1] && n[1] == n[2] {
		fmt.Fprint(out, "Yes")

	} else {
		fmt.Fprint(out, "No")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
