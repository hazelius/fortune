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
	result := binarySearch(n, 1, 999999)
	return result
}

func double(n int) int {
	nstr := strconv.Itoa(n)
	nn, _ := strconv.Atoi(nstr + nstr)
	return nn
}

func binarySearch(n, min, max int) int {
	mid := min + (max-min)/2
	m := double(mid)
	// fmt.Printf("n:%v, m:%v, min:%v, max:%v, mid:%v\n", n, m, min, max, mid)

	if m == n {
		return mid
	}

	if n > m {
		if mid == max {
			return mid
		}
		return binarySearch(n, mid+1, max)
	} else {
		if mid == min {
			return mid - 1
		}
		return binarySearch(n, min, mid-1)
	}
}

func pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	fmt.Println(run(n))
}
