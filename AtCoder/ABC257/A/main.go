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

	n, x := readInt(), readInt()

	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	i := x / n
	if x%n > 0 {
		i++
	}

	return alpha[i-1 : i]
}

func main() {
	fmt.Println(run(os.Stdin))
}
