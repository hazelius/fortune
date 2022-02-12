// https://atcoder.jp/contests/abc216/tasks/abc216_e
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

	low, high := 0, int(1e9*2+1)
	for low+1 < high {
		mid := low + (high-low)/2
		sum := 0
		for _, a := range as {
			if a > mid {
				sum += a - mid
			}
		}
		if k < sum {
			low = mid
		} else {
			high = mid
		}
	}

	ans := 0
	sum := 0
	for _, a := range as {
		if a > low {
			s := a - low
			ans += low * s
			ans += s * (s + 1) / 2
			sum += a - low
		}
	}

	if sum > k {
		p := sum - k
		ans -= p * high
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
