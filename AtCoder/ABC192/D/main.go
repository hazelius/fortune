package main

import (
	"fmt"
	"strconv"
)

func max(y string) int {
	result := 0
	for i := range y {
		j, _ := strconv.Atoi(y[i : i+1])
		if result < j {
			result = j
		}
	}
	return result
}

func nshinho(x string, n, m uint64) uint64 {
	var result uint64
	for i := range x {
		v, _ := strconv.ParseUint(x[i:i+1], 10, 64)
		for j := i; j < len(x)-1; j++ {
			v *= n
			if v > m {
				return v
			}
		}
		result += v
	}
	return result
}

func run(y string, m uint64) string {
	if len(y) == 1 {
		x, _ := strconv.ParseUint(y, 10, 64)
		if x <= m {
			return "1"
		}
		return "0"
	}

	x := max(y)
	var i, r uint64
	min := uint64(x + 1)
	pre := min
	for i = min; r < m; i += 10000 {
		r = nshinho(y, i, m)
		if r >= m {
			break
		}
		pre = i
	}

	r = 0

	for i = pre; r < m; i++ {
		r = nshinho(y, i, m)
		if r >= m {
			break
		}
	}

	if r == m {
		i++
	}
	return fmt.Sprint(i - min)
}

func main() {
	var y string
	var m uint64

	fmt.Scanf("%s", &y)
	fmt.Scanf("%d", &m)
	fmt.Println(run(y, m))
}
