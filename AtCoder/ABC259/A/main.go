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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	_, m, x, t, d := readInt(), readInt(), readInt(), readInt(), readInt()
	if m >= x {
		return t
	}
	return t - (x-m)*d
}

func main() {
	fmt.Println(run(os.Stdin))
}
