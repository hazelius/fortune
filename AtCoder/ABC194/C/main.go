package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r []int) string {
	m := make(map[int]int, 0)
	for _, v := range r {
		m[v]++
	}

	var result int
	for i := -200; i < 200; i++ {
		a, ok := m[i]
		if !ok {
			continue
		}
		for j := i + 1; j <= 200; j++ {
			b, ok := m[j]
			if !ok {
				continue
			}
			v := j - i
			result += a * b * v * v
		}
	}

	return fmt.Sprint(result)
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()

	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = readInt()
	}
	result := run(r)

	fmt.Println(result)
}
