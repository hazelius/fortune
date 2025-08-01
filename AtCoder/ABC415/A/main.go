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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	x := readInt()

	for _, a := range as {
		if x == a {
			fmt.Fprint(out, "Yes")
			return
		}
	}
	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
