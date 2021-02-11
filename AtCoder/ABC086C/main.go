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
func isTrue(t, x, y int64) bool {
	result := t - (abs(x) + abs(y))
	if result < 0 {
		return false
	}
	if result%2 > 0 {
		return false
	}
	return true
}

func run(txy [][3]int64) string {
	var start [3]int64
	for _, v := range txy {
		if !isTrue(v[0]-start[0], v[1]-start[1], v[2]-start[2]) {
			return "No"
		}
		start = v
	}

	return "Yes"
}

func main() {
	var n int64
	fmt.Scanf("%d", &n)
	txy := make([][3]int64, n)
	for i := range txy {
		fmt.Scanf("%d %d %d", &txy[i][0], &txy[i][1], &txy[i][2])
	}
	fmt.Println(run(txy))
}
