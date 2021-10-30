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

	n := readInt()
	ps := make(map[int]int, n)
	for i := 0; i < n; i++ {
		a, b := readInt(), readInt()
		ps[a]++
		ps[b]++
	}

	for _, v := range ps {
		if v == n-1 {
			return "Yes"
		}
	}
	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
