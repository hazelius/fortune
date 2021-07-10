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

	n, x := readInt(), readInt()
	s := 0
	for i := 1; i <= n; i++ {
		a := readInt()
		if i%2 == 0 {
			a--
		}
		s += a
	}
	if s <= x {
		return "Yes"
	}
	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
