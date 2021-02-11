package main

import "fmt"

func run(v, t, s, d int) string {
	if v*t <= d && d <= v*s {
		return "No"
	}
	return "Yes"
}

func main() {
	var v, t, s, d int
	fmt.Scanf("%d %d %d %d", &v, &t, &s, &d)
	fmt.Println(run(v, t, s, d))
}
