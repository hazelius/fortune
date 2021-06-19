package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var flags map[int]bool

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, as []int) int {
	m := make(map[int]map[int]bool)
	for i := 0; i < n/2; i++ {
		a := as[i]
		b := as[n-i-1]
		if a != b {
			if _, ok := m[a]; !ok {
				m[a] = make(map[int]bool)
			}
			m[a][b] = true
			if _, ok := m[b]; !ok {
				m[b] = make(map[int]bool)
			}
			m[b][a] = true
		}
	}

	ans := 0
	flags = make(map[int]bool)
	for k := range m {
		if _, ok := flags[k]; !ok {
			ans += dfs(k, m) - 1
		}
	}
	return ans
}

func dfs(val int, m map[int]map[int]bool) int {
	if _, ok := flags[val]; ok {
		return 0
	}

	flags[val] = true

	ans := 1
	for k := range m[val] {
		ans += dfs(k, m)
	}
	return ans
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
