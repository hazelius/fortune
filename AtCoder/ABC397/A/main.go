package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readFloat() float64 {
	sc.Scan()
	i, _ := strconv.ParseFloat(sc.Text(), 64)
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	x := readFloat()
	x *= 10
	if x >= 380 {
		fmt.Fprint(out, 1)
	} else if x >= 375 {
		fmt.Fprint(out, 2)
	} else {
		fmt.Fprint(out, 3)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
