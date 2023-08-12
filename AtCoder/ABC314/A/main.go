package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	p := "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"

	n := readInt()
	fmt.Fprint(out, p[:n+2])
}

func main() {
	run(os.Stdin, os.Stdout)
}
