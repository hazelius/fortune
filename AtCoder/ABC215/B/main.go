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
	ans := 0
	b := 1
	for {
		b *= 2
		if b > n {
			break
		}
		ans++
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
