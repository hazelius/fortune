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

	x := int(1e18)
	j := int(1e6)
	for i := 0; i < 1e6; i++ {
		for ; i <= j && j >= 0; j-- {
			c := f(i, j)
			if c < n {
				break
			}
			if c < x {
				x = c
			}
		}
	}

	return x
}

func f(a, b int) int {
	return a*a*a + a*a*b + a*b*b + b*b*b
}

func main() {
	fmt.Println(run(os.Stdin))
}
