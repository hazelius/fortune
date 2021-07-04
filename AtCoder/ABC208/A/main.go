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

	a, b := readInt(), readInt()
	if a <= b && b <= a*6 {
		return "Yes"
	}
	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
