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
	odds := []int{-1, -1}
	evens := []int{-1, -1}

	for i := 0; i < n; i++ {
		a := readInt()
		if a%2 > 0 {
			if odds[0] < a {
				odds[1] = odds[0]
				odds[0] = a
			} else if odds[1] < a {
				odds[1] = a
			}
		} else {
			if evens[0] < a {
				evens[1] = evens[0]
				evens[0] = a
			} else if evens[1] < a {
				evens[1] = a
			}
		}
	}
	ans := -1
	if odds[0] >= 0 && odds[1] >= 0 {
		ans = odds[0] + odds[1]
	}
	if evens[0] >= 0 && evens[1] >= 0 {
		maxe := evens[0] + evens[1]
		if maxe > ans {
			ans = maxe
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
