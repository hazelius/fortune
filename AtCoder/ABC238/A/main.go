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

	n := readInt()
	nn := n * n
	n2 := 1
	for i := 0; i < n; i++ {
		n2 *= 2
		if nn < n2 {
			return "Yes"
		}
	}

	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
