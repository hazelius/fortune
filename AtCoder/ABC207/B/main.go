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
	a, b, c, d := readInt(), readInt(), readInt(), readInt()

	if b >= c*d {
		return -1
	}

	if a == 0 {
		return 0
	}

	i, red := 0, 0
	for ; a > red*d; i++ {
		a += b
		red += c
	}
	return i
}

func main() {
	fmt.Println(run(os.Stdin))
}
