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

	a, b, c := readInt(), readInt(), readInt()
	if b > a && b > c {
		return "No"
	}
	if b < a && b < c {
		return "No"
	}
	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
