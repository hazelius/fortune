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

	a, b, c := readInt(), readInt(), readInt()

	ans := -1
	for i := 0; i <= b; i += c {
		if a <= i && i <= b {
			ans = i
			break
		}
		if b < i {
			break
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
