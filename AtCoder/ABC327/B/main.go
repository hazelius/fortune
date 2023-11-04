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

	b := readInt()

	for i := 1; i < 17; i++ {
		a := 1
		for j := 0; j < i; j++ {
			a *= i
		}
		if a == b {
			fmt.Fprint(out, i)
			return
		}
		if a > b {
			break
		}
	}

	fmt.Fprint(out, -1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
