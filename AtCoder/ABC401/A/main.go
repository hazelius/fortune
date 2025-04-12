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

	n := readInt()
	if n >= 200 && n < 300 {
		fmt.Fprint(out, "Success")
		return
	}
	fmt.Fprint(out, "Failure")
}

func main() {
	run(os.Stdin, os.Stdout)
}
