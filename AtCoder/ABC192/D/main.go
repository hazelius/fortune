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

func run(y string, m int) string {
	if len(y) == 1 {
		x, _ := strconv.Atoi(y)
		if x <= m {
			return "1"
		}
		return "0"
	}

	min := max(y) + 1
	max := binarySearch(y, m, min, int(1e18+1))
	return fmt.Sprint(max - min + 1)
}

func formatInt(i, base int) []int {
	ret := make([]int, 0)
	for cnt := 0; ; cnt++ {
		ret = append([]int{i % base}, ret...)
		i = i / base
		if i < base {
			if i != 0 {
				ret = append([]int{i}, ret...)
			}
			break
		}
	}
	return ret
}

func binarySearch(x string, m, min, max int) int {
	mid := min + (max-min)/2
	ms := formatInt(m, mid)
	// fmt.Printf("x:%v, m:%v, min:%v, max:%v, mid:%v, ms:%v\n", x, m, min, max, mid, ms)

	res := 0
	if len(x) < len(ms) {
		res = 1
	} else if len(x) > len(ms) {
		res = -1
	} else {
		for i, mv := range ms {
			xi, _ := strconv.Atoi(x[i : i+1])
			if xi < mv {
				res = 1
				break
			} else if xi > mv {
				res = -1
				break
			}
		}
	}

	if res == 0 {
		return mid
	}

	if res > 0 {
		if mid == max {
			return mid
		}
		return binarySearch(x, m, mid+1, max)
	} else {
		if mid == min {
			return mid - 1
		}
		return binarySearch(x, m, min, mid-1)
	}
}

func main() {
	var y string
	var m int

	fmt.Scanf("%s", &y)
	fmt.Scanf("%d", &m)
	fmt.Println(run(y, m))
}
