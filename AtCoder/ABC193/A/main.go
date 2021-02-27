package main

import (
	"fmt"
)

func run(a, b float64) string {
	r := 100 * (a - b)
	return fmt.Sprintf("%g", r/a)
}

func main() {
	var a, b float64
	fmt.Scanf("%g %g", &a, &b)
	fmt.Println(run(a, b))
}
