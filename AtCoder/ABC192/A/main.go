package main

import (
	"fmt"
)

func run(x int64) string {
	a := x % 100
	r := 100 - a
	return fmt.Sprintf("%d", r)
}

func main() {
	var x int64
	fmt.Scanf("%d %d", &x)
	fmt.Println(run(x))
}
