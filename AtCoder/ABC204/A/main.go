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

func run(x, y int) int {
	if x == y {
		return x
	}
	ret := (x + 1) % 3
	if ret == y {
		ret = (ret + 1) % 3
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	x, y := readInt(), readInt()
	fmt.Println(run(x, y))
}
