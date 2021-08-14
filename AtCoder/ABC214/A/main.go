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

	n := readInt()
	if n < 126 {
		return 4
	} else if n < 212 {
		return 6
	}
	return 8
}

func main() {
	fmt.Println(run(os.Stdin))
}
