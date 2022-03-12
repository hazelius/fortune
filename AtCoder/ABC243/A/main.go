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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	v, a, b, c := readInt(), readInt(), readInt(), readInt()
	s := a + b + c
	amari := v % s
	if amari < a {
		return "F"
	}
	amari -= a
	if amari < b {
		return "M"
	}
	return "T"
}

func main() {
	fmt.Println(run(os.Stdin))
}
