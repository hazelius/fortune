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

	n, k := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	low, high := 0, 300001
	for low+1 < high {
		mid := low + (high-low)/2
		if isOK(mid, k, as) {
			low = mid
		} else {
			high = mid
		}
	}

	return low
}

func isOK(gdc, k int, as []int) bool {
	ret := 0
	for _, a := range as {
		amari := a % gdc
		if amari > 0 {
			ret += gdc - amari
		}
		if ret > k {
			return false
		}
	}

	return ret <= k
}

func main() {
	fmt.Println(run(os.Stdin))
}
