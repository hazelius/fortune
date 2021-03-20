package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func run(a, b, c, d int) int {
	m1 := max(a, b) - min(c, d)
	m2 := min(a, b) - max(c, d)

	return max(m1, m2)
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b := readInt(), readInt()
	c, d := readInt(), readInt()

	fmt.Println(run(a, b, c, d))
}
