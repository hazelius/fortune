package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n [][]int, m []int, q [][]int) []int {
	result := make([]int, len(q))

	sort.Slice(n, func(i, j int) bool {
		return n[i][1] > n[j][1]
	})
	for i, qv := range q {
		arg := make([]int, len(m)-(qv[1]-qv[0]+1))
		j := 0
		for mi, v := range m {
			if mi >= qv[0]-1 && mi <= qv[1]-1 {
				continue
			}
			arg[j] = v
			j++
		}

		result[i] = qresult(n, arg)
	}

	return result
}

func qresult(n [][]int, m []int) int {
	sets := make([]int, 0, len(m))
	sort.Ints(m)

	for _, mv := range m {
		for i, nv := range n {
			if settable(mv, nv[0], i, sets) {
				sets = append(sets, i)
				break
			}
		}
	}

	var result int
	for _, i := range sets {
		result += n[i][1]
	}
	return result
}

func settable(m, n, i int, sets []int) bool {
	for _, d := range sets {
		if i == d {
			return false
		}
	}

	if n <= m {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m, q := readInt(), readInt(), readInt()

	ns := make([][]int, n)
	for i := 0; i < n; i++ {
		ns[i] = []int{readInt(), readInt()}
	}

	ms := make([]int, m)
	for i := 0; i < m; i++ {
		ms[i] = readInt()
	}

	qs := make([][]int, q)
	for i := 0; i < q; i++ {
		qs[i] = []int{readInt(), readInt()}
	}

	result := run(ns, ms, qs)
	for _, r := range result {
		fmt.Println(r)
	}
}
