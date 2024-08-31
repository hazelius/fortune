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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ans := 0
	prel := -1
	prer := -1
	for i := 0; i < n; i++ {
		a := readInt()
		s := readString()
		if s == "L" {
			if prel >= 0 {
				ans += abs(prel - a)
			}
			prel = a
		} else {
			if prer >= 0 {
				ans += abs(prer - a)
			}
			prer = a
		}
	}
	fmt.Fprint(out, ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
