package main

import (
	"fmt"
	"strings"
)

var items = []string{"dream", "dreamer", "erase", "eraser"}

func erase(s string) (bool, string) {
	for _, v := range items {
		if strings.HasSuffix(s, v) {
			s = s[:len(s)-len(v)]
			return true, s
		}
	}
	return false, s
}

func run(s string) string {
	var result bool
	for len(s) > 0 {
		result, s = erase(s)
		if !result {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	var s string
	fmt.Scanf("%v", &s)
	fmt.Println(run(s))
}
