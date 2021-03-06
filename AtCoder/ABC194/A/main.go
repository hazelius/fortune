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

func run(a, b int) string {
	s := a + b
	if s >= 15 && b >= 8 {
		return "1"
	} else if s >= 10 && b >= 3 {
		return "2"
	} else if s >= 3 {
		return "3"
	}

	return "4"
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b := readInt(), readInt()
	result := run(a, b)

	fmt.Println(result)
}
