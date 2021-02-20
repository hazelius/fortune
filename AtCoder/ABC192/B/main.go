package main

import (
	"fmt"
	"strings"
)

func run(s string) string {
	for i := range s {
		c := s[i : i+1]
		if i%2 == 0 {
			if c == strings.ToUpper(c) {
				return "No"
			}
		} else {
			if c == strings.ToLower(c) {
				return "No"
			}
		}
	}
	return "Yes"
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(run(s))
}
