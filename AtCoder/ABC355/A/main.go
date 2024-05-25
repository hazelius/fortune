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
		fmt.Fprint(out, -1)
		return
	}
	for i := 1; i <= 3; i++ {
		if i != a && i != b {
			fmt.Fprint(out, i)
			return
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
