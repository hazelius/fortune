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

	a, b := readInt(), readInt()
	ans := b - a + 1
	if ans < 0 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
