package main

import (
	"fmt"
)

func run(n uint64) string {
	var s, i uint64
	m := make(map[uint64]bool, 0)

	for i = 2; ; i++ {
		if m[i] {
			continue
		}

		j := i
		flg := false
		for {
			j = j * i
			if j <= n {
				m[j] = true
				s++
				flg = true
			} else {
				break
			}
		}
		if !flg {
			break
		}
	}

	r := n - s
	return fmt.Sprintf("%v", r)
}

func main() {
	var n uint64
	fmt.Scanf("%d", &n)
	fmt.Println(run(n))
}
