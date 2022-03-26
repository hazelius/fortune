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

	n := readInt()
	amap := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a] = true
	}
	for i := 0; i < 2001; i++ {
		if !amap[i] {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(run(os.Stdin))
}
