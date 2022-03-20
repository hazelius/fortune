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

func run(r io.Reader) {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]bool, 2*n+2)
	as[0] = true
	idx := 1
	as[idx] = true
	fmt.Println(idx)
	a := readInt()
	for a != 0 {
		as[a] = true
		for as[idx] {
			idx++
		}
		as[idx] = true
		fmt.Println(idx)
		a = readInt()
	}
}

func main() {
	run(os.Stdin)
}
