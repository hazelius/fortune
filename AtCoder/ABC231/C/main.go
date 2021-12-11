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
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		x := readInt()

		low, high := 0, n
		for low+1 < high {
			mid := low + (high-low)/2
			if x >= as[mid] {
				low = mid
			} else {
				high = mid
			}
		}
		ans[i] = n - low - 1
		if as[low] >= x {
			ans[i]++
		}
	}
	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
