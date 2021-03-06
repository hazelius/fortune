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

func run(r [][]int) string {
	mina, minb := -1, -1
	mina2, minb2 := -1, -1
	for i, v := range r {
		if mina < 0 || v[0] <= r[mina][0] {
			mina2 = mina
			mina = i
		}
		if minb < 0 || v[1] <= r[minb][1] {
			minb2 = minb
			minb = i
		}
	}
	result := max(r[mina][0], r[minb][1])

	if mina == minb {
		l := r[mina][0] + r[minb][1]
		if mina2 < 0 && minb2 < 0 {
			return fmt.Sprint(l)
		}

		var min2 int
		if mina2 < 0 {
			min2 = r[minb2][1]
		} else if minb2 < 0 {
			min2 = r[mina2][0]
		} else {
			min2 = min(r[mina2][0], r[minb2][1])
		}

		if min2 < l {
			if min2 > result {
				result = min2
			}
		} else {
			result = l
		}
	}

	return fmt.Sprint(result)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()

	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, 2)
		r[i][0], r[i][1] = readInt(), readInt()
	}
	result := run(r)

	fmt.Println(result)
}
