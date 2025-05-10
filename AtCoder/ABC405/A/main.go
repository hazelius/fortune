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

	r, x := readInt(), readInt()
	if x == 1 {
		if r >= 1600 && r < 3000 {
			fmt.Fprint(out, "Yes")
			return
		}
	} else {
		if r >= 1200 && r < 2400 {
			fmt.Fprint(out, "Yes")
			return
		}
	}
	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
