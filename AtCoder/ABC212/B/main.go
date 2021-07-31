package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	x := readString()
	flg := false
	for i := 0; i < len(x)-1; i++ {
		if x[i] != x[i+1] {
			flg = true
		}
	}
	if !flg {
		return "Weak"
	}

	flg = false
	for i := 0; i < len(x)-1; i++ {
		x1, _ := strconv.Atoi(x[i : i+1])
		x2, _ := strconv.Atoi(x[i+1 : i+2])
		next := x1 + 1
		if x1 == 9 {
			next = 0
		}
		if x2 != next {
			flg = true
		}
	}
	if flg {
		return "Strong"
	}

	return "Weak"
}

func main() {
	fmt.Println(run(os.Stdin))
}
