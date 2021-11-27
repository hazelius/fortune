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

	a, b := readInt(), readInt()
	for a > 0 && b > 0 {
		ad := a % 10
		bd := b % 10
		if ad+bd > 9 {
			return "Hard"
		}
		a /= 10
		b /= 10
	}
	return "Easy"
}

func main() {
	fmt.Println(run(os.Stdin))
}
