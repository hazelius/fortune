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

	x := readInt()
	for i := 2; i <= x; i++ {
		x /= i
		if x == 1 {
			fmt.Fprint(out, i)
			return
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
