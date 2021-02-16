package main

import (
	"fmt"
)

func abs(n int64) int64 {
	if n < 0 {
		return n * -1
	}
	return n
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func run(b, c int64) string {
	maxA := b + ((c - 2) / 2)
	minA := b - (c / 2)
	maxB := b*-1 + ((c - 1) / 2)
	minB := b*-1 - ((c - 1) / 2)

	r := abs(maxA-minA) + abs(maxB-minB) + 2
	d := min(maxA, maxB) - max(minA, minB) + 1
	if d > 0 {
		r -= d
	}

	return fmt.Sprintf("%v", r)
}

func main() {
	var b, c int64
	fmt.Scanf("%d %d", &b, &c)
	fmt.Println(run(b, c))
}
