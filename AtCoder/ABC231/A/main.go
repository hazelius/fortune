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

func run(r io.Reader) float64 {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	d := readInt()
	return float64(d) / 100
}

func main() {
	fmt.Println(run(os.Stdin))
}
