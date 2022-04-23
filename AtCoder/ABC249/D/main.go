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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	amap := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a]++
	}
	as := make([]int, len(amap))
	i := 0
	for k := range amap {
		as[i] = k
		i++
	}

	sort.Ints(as)

	ans := 0
	for _, a1 := range as {
		for _, a2 := range as {
			if a1*a2 > as[len(as)-1] {
				break
			}
			a3 := a1 * a2
			a3c, ok := amap[a3]
			if !ok {
				continue
			}
			ans += amap[a1] * amap[a2] * a3c
		}
	}

	// sort.Sort(sort.Reverse(sort.IntSlice(as)))
	// for i := 0; i < len(as)-1; i++ {
	// 	for j := i; j < len(as); j++ {
	// 		if as[i] < as[j]*2 || as[i] > as[j]*as[j] {
	// 			continue
	// 		}

	// 		for k := j; k < len(as); k++ {
	// 			if as[i]%as[k] == 0 && as[i]/as[k] == as[j] {
	// 				ans += amap[as[i]] * amap[as[j]] * amap[as[k]]
	// 			}
	// 			if j != k {
	// 				if as[i]%as[j] == 0 && as[i]/as[j] == as[k] {
	// 					ans += amap[as[i]] * amap[as[j]] * amap[as[k]]
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// if cnt, ok := amap[1]; ok {
	// 	if cnt >= 3 {
	// 		ans += cnt * cnt * cnt
	// 	}
	// 	for k, c := range amap {
	// 		if c == 1 || k == 1 {
	// 			continue
	// 		}
	// 		ans += c * c * cnt
	// 	}
	// }

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
