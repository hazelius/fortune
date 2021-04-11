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

func run(n int) int {
	return n - 1
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	fmt.Println(run(n))
}
