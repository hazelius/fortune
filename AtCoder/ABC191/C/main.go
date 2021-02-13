package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a := nextLine()
	b := strings.Split(a, " ")
	fmt.Println(b)
}
