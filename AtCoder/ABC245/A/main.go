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

	a, b, c, d := readInt(), readInt(), readInt(), readInt()
	if a > c {
		return "Aoki"
	} else if a < c {
		return "Takahashi"
	}
	if b > d {
		return "Aoki"
	} else if b < d {
		return "Takahashi"
	}
	return "Takahashi"
}

func main() {
	fmt.Println(run(os.Stdin))
}
