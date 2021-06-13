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

func run(n, q int, as, ks []int) []int {
	ans := make([]int, q)
	for i, k := range ks {
		ans[i] = f(k, as)
	}
	return ans
}

func f(k int, as []int) int {
	low := 0
	high := len(as)
	for low+1 < high {
		mid := low + (high-low)/2
		if as[mid] > k+mid {
			high = mid
		} else if as[mid] < k+mid {
			low = mid
		} else {
			break
		}
	}

	idx := k + low
	for i := low; i < len(as); i++ {
		if as[i] <= idx {
			idx++
		} else {
			break
		}
	}
	return idx
}

func main() {
	sc.Split(bufio.ScanWords)
	n, q := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	ks := make([]int, q)
	for i := range ks {
		ks[i] = readInt()
	}
	ans := run(n, q, as, ks)
	for _, v := range ans {
		fmt.Println(v)
	}
}
