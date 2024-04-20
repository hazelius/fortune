package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
	n, _ := strconv.Atoi(s[3:])
	if n >= 350 || n < 1 || n == 316 {
		fmt.Fprint(out, "No")
	} else {
		fmt.Fprint(out, "Yes")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
