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

	h := readInt()
	v := 1
	i := 0
	for ; h >= 0; i++ {
		h -= v
		v *= 2
	}
	fmt.Fprint(out, i)
}

func main() {
	run(os.Stdin, os.Stdout)
}
