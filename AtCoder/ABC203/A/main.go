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

func run(a, b, c int) int {
	if a == b {
		return c
	} else if a == c {
		return b
	} else if b == c {
		return a
	}
	return 0
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b, c := readInt(), readInt(), readInt()
	fmt.Println(run(a, b, c))
}
