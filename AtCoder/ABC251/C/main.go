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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	poems := make(map[string]bool)
	high := 0
	ans := 0
	for i := 0; i < n; i++ {
		s := readString()
		t := readInt()
		if !poems[s] {
			poems[s] = true
			if high < t {
				high = t
				ans = i
			}
		}
	}
	return ans + 1
}

func main() {
	fmt.Println(run(os.Stdin))
}
