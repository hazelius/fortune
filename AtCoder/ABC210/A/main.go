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

	n, a, x, y := readInt(), readInt(), readInt(), readInt()
	if n <= a {
		return x * n
	}
	ans := x * a
	ans += y * (n - a)
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
