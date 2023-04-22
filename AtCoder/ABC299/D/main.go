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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	low, high := 0, n
	for low+1 < high {
		mid := low + (high-low)/2
		fmt.Fprintf(out, "? %v\n", mid+1)
		s := readInt()

		if s == 0 {
			low = mid
		} else {
			high = mid
		}
	}

	fmt.Fprintf(out, "! %v\n", low+1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
