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

func run(m, h int) string {
	if h%m == 0 {
		return "Yes"
	}

	return "No"
}

func main() {
	sc.Split(bufio.ScanWords)
	m, h := readInt(), readInt()
	result := run(m, h)

	fmt.Println(result)
}
