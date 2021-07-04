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

	p := readInt()
	coins := make([]int, 11)
	for i := range coins {
		if i == 0 {
			coins[i] = 1
		} else {
			coins[i] = coins[i-1] * i
		}
	}
	ans := 0
	for i := 10; i >= 1; i-- {
		ans += p / coins[i]
		p %= coins[i]
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
