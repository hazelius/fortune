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
	amap := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a]++
	}
	return len(amap)
}

func main() {
	fmt.Println(run(os.Stdin))
}
