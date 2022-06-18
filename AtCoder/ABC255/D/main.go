package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

	n, q := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	sort.Ints(as)
	mina, maxa := as[0], as[n-1]

	minSum := 0
	dsmin := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		minSum += as[i] - mina
		dsmin[i] = minSum
	}

	maxSum := 0
	dsmax := make([]int, n)
	for i, a := range as {
		maxSum += maxa - a
		dsmax[i] = maxSum
	}

	ans := make([]int, q)
	for i := range ans {
		x := readInt()
		if x <= mina {
			ans[i] = minSum + (mina-x)*n
			continue
		}
		if x >= maxa {
			ans[i] = maxSum + (x-maxa)*n
			continue
		}

		low, high := 0, n
		for low+1 < high {
			mid := low + (high-low)/2
			if x == as[mid] {
				low = mid
				break
			}
			if x > as[mid] {
				low = mid
			} else {
				high = mid
			}
		}
		sum := dsmax[low] - (low+1)*(maxa-x)
		sum += dsmin[low+1] - (n-(low+1))*(x-mina)
		ans[i] = sum
	}
	anss := fmt.Sprintf("%v", ans)
	return anss[1 : len(anss)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
