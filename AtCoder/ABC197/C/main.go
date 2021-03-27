package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(arr []int) int {
	r := math.MaxInt32

	// bit全検索
	for bit := 0; bit < (1<<len(arr) - 1); bit++ {
		xor, or := 0, 0
		for i := 0; i < len(arr); i++ {
			if bit>>i&1 > 0 {
				// fmt.Printf("%b, %b\n", or, xor)
				xor ^= or
				or = 0
			}

			or |= arr[i]
		}
		xor ^= or

		if r > xor {
			r = xor
		}
	}
	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	fmt.Println(run(arr))
}
