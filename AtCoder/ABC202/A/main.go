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
	return (7 - a) + (7 - b) + (7 - c)
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b, c := readInt(), readInt(), readInt()
	fmt.Println(run(a, b, c))
}
