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
	if a == 0 {
		return "Silver"
	} else if b == 0 {
		return "Gold"
	}
	return "Alloy"
}

func main() {
	fmt.Println(run(os.Stdin))
}
