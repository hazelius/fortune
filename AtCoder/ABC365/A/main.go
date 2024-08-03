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

	y := readInt()
	if y%4 > 0 {
		fmt.Fprint(out, 365)
	} else if y%100 > 0 {
		fmt.Fprint(out, 366)
	} else if y%400 > 0 {
		fmt.Fprint(out, 365)
	} else {
		fmt.Fprint(out, 366)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
