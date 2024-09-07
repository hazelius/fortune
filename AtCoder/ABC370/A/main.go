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

	l, r := readInt(), readInt()
	if l == 1 && r == 0 {
		fmt.Fprint(out, "Yes")

	} else if l == 0 && r == 1 {
		fmt.Fprint(out, "No")
	} else {
		fmt.Fprint(out, "Invalid")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
