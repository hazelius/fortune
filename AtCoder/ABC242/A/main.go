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

func run(r io.Reader) float64 {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	a, b, c, x := readInt(), readInt(), readInt(), readInt()
	if x <= a {
		return 1
	}
	if x > b {
		return 0
	}
	return float64(c) / float64(b-a)
}

func main() {
	fmt.Println(run(os.Stdin))
}
