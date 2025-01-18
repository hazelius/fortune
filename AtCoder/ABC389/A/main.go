package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

	s := readString()
	ss := strings.Split(s, "x")
	a, _ := strconv.Atoi(ss[0])
	b, _ := strconv.Atoi(ss[1])
	fmt.Fprint(out, a*b)
}

func main() {
	run(os.Stdin, os.Stdout)
}
