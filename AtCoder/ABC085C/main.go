package main

import "fmt"

func calcSum(a, b, c int64) int64 {
	return a*10000 + b*5000 + c*1000
}

func run(n, y int64) string {
	var z int64

	for a := z; a <= n; a++ {
		for b := z; b <= n-a; b++ {
			c := n - a - b
			if calcSum(a, b, c) == y {
				return fmt.Sprintf("%d %d %d", a, b, c)
			}
		}
	}
	return "-1 -1 -1"
}

func main() {
	var n, y int64
	fmt.Scanf("%d %d", &n, &y)
	fmt.Println(run(n, y))
}
