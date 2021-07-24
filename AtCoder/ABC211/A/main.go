package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readFloat() float64 {
	sc.Scan()
	i, _ := strconv.ParseFloat(sc.Text(), 64)
	return i
}

func run(r io.Reader) float64 {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	a, b := readFloat(), readFloat()

	return (a + 2*b) / 3
}

func main() {
	fmt.Println(run(os.Stdin))
}
