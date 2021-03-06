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

func run(r []int, m int) string {
	mp := make(map[int]int, 0)
	for i := 0; i < m; i++ {
		mp[r[i]]++
	}
	mm := mex(r[:m])
	if mm == 0 {
		return "0"
	}
	// fmt.Printf("mm:%v, mp:%v\n", mm, mp)

	last := len(r) - m
	for i := 1; i <= last; i++ {
		drop := r[i-1]
		mp[drop]--
		push := r[i+m-1]
		mp[push]++
		// fmt.Printf("drop:%v, push:%v, mm:%v, mp:%v\n", drop, push, mm, mp)

		if mp[drop] == 0 && drop < mm {
			mm = drop
		}
	}

	return fmt.Sprint(mm)
}

func mex(r []int) int {
	m := make(map[int]bool, len(r))
	max := 0
	for _, v := range r {
		m[v] = true
		if max < v {
			max = v
		}
	}
	for i := 0; i < max; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return max + 1
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := readInt(), readInt()
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = readInt()
	}
	fmt.Println(run(r, m))
}
