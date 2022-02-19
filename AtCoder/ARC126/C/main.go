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
	maxA := 0
	for i := range as {
		a := readInt()
		as[i] = a
		if maxA < a {
			maxA = a
		}
	}

	tc := 0
	for _, a := range as {
		tc += maxA - a
	}
	if tc <= k {
		k -= tc
		return maxA + (k / n)
	}

	cnt := make([]int, 2*maxA)
	sum := make([]int, 2*maxA)

	for _, a := range as {
		cnt[a]++
		sum[a] += a
	}

	for i := 1; i < 2*maxA; i++ {
		cnt[i] += cnt[i-1]
		sum[i] += sum[i-1]
	}

	ans := 1
	for x := 2; x < maxA; x++ {
		bound := 0
		for i := 1; i <= 1+(maxA-1)/x; i++ {
			bound += i*x*(cnt[i*x]-cnt[(i-1)*x]) - (sum[i*x] - sum[(i-1)*x])
		}
		if bound <= k {
			ans = x
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
