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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	asmap := make(map[int][]int)
	for i := 0; i < n; i++ {
		a := readInt()
		amap, ok := asmap[a]
		if ok {
			asmap[a] = append(amap, i+1)
		} else {
			asmap[a] = []int{i + 1}
		}
	}

	q := readInt()
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		l, r, k := readInt(), readInt(), readInt()
		amap, ok := asmap[k]
		if !ok {
			continue
		}
		if len(amap) == 1 {
			if l <= amap[0] && amap[0] <= r {
				ans[i] = 1
			}
			continue
		}
		li := f(l, amap)
		ri := f(r+1, amap)
		if amap[li] < l {
			li++
		}
		if amap[ri] <= r {
			ri++
		}

		ans[i] = ri - li
	}
	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func f(t int, as []int) int {
	low, high := 0, int(1e9+1)
	for low+1 < high {
		mid := low + (high-low)/2
		if len(as) > mid && as[mid] < t {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	fmt.Println(run(os.Stdin))
}
