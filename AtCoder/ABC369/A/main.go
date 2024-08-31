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

	a, b := readInt(), readInt()
	if a == b {
		fmt.Fprint(out, 1)
	} else if (a-b)%2 != 0 {
		fmt.Fprint(out, 2)
	} else {
		fmt.Fprint(out, 3)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
