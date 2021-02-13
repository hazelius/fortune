package main

import (
	"fmt"
)

func run(x int, a []int) string {
	result := make([]int, 0, len(a))
	for _, v := range a {
		if v != x {
			result = append(result, v)
		}
	}
	re := fmt.Sprintf("%v", result)
	return re[1 : len(re)-1]
}

func main() {
	var n, x int
	fmt.Scanf("%d %d", &n, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Println(run(x, a))
}
