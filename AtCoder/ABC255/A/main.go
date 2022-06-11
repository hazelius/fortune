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

	rc, c := readInt(), readInt()
	as := make([][]int, 2)
	for i := range as {
		as[i] = []int{readInt(), readInt()}
	}
	return as[rc-1][c-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
