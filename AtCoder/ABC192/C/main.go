package main

import (
	"fmt"
	"sort"
)

func g1(x int64, flg bool) int64 {
	arr := make([]int64, 0)
	for x >= 10 {
		a := x % 10
		arr = append(arr, a)
		x = x / 10
	}
	arr = append(arr, x)
	if flg {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	} else {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})

	}

	var result, j int64
	j = 1
	for _, v := range arr {
		result += v * j
		j *= 10
	}

	return result
}

func f(x int64) int64 {
	return g1(x, true) - g1(x, false)
}

func run(n, k int64) string {
	var i int64
	for i = 0; i < k; i++ {
		n = f(n)
	}
	return fmt.Sprintf("%v", n)
}

func main() {
	var n, k int64
	fmt.Scanf("%d %d", &n, &k)
	fmt.Println(run(n, k))
}
