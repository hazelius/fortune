package main

import (
	"fmt"
)

func k(n int64) int64 {
	return n * (n + 1) / 2
}

func cased(l, r int64) int64 {
	for j := l; j <= r; j++ {
		s := r - j
		if s >= l {
			return k((s - l + 1))
		}
	}
	return 0
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var l, r int64
		fmt.Scanf("%d %d", &l, &r)
		fmt.Println(cased(l, r))
	}
}
