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
	as := make([]int, n)
	sum := 0
	for i := range as {
		a := readInt()
		sum += a
		as[i] = sum
	}
	x := readInt()

	ans := 0
	ans = (x / sum) * n
	x = x % sum

	for i, a := range as {
		if x < a {
			ans += i + 1
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
