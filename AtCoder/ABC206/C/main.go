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

func run(n int, as []int) int {
	m := make(map[int]int)
	for _, v := range as {
		cnt, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] = cnt + 1
		}
	}

	dup := 0
	for _, cnt := range m {
		if cnt > 1 {
			dup += cnt * (cnt - 1)
		}
	}
	return (n*(n-1) - dup) / 2
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, as))
}
