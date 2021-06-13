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

func run(a, b, c int) string {
	if c%2 == 0 {
		d := abs(a) - abs(b)
		if d == 0 {
			return "="
		} else if d > 0 {
			return ">"
		}
		return "<"
	}

	if a == b {
		return "="
	} else if a > b {
		return ">"
	}
	return "<"
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b, c := readInt(), readInt(), readInt()
	fmt.Println(run(a, b, c))
}
