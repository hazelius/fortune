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

	a, b, k := readInt(), readInt(), readInt()
	ans := 0
	for i := a; i < b; i *= k {
		ans++
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
