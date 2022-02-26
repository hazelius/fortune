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

	as := make([]int, 10)
	for i := range as {
		as[i] = readInt()
	}
	return as[as[as[0]]]
}

func main() {
	fmt.Println(run(os.Stdin))
}
