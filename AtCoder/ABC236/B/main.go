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

	n := readInt()
	cards := make([]int, n)
	for i := 0; i < n*4-1; i++ {
		a := readInt()
		cards[a-1]++
	}

	for i, v := range cards {
		if v < 4 {
			return i + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(run(os.Stdin))
}
