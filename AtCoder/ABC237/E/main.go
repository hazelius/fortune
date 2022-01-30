package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	hs := make([]int, n+1)
	for i := 0; i < n; i++ {
		hs[i+1] = readInt()
	}
	uvmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		u, v := readInt(), readInt()
		uvmap[u] = append(uvmap[u], v)
		uvmap[v] = append(uvmap[v], u)
	}

	queue := []int{1}
	emap := make(map[int]int)
	emap[1] = 0
	used := make(map[int]bool)
	used[1] = true
	mx := 0

	for {
		if len(queue) == 0 {
			break
		}
		fm := queue[0]
		queue = queue[1:]
		qs, ok := uvmap[fm]
		if !ok {
			continue
		}
		for _, to := range qs {
			e := hs[fm] - hs[to]
			if e < 0 {
				e *= 2
			}

			sume := emap[fm] + e
			mx = max(mx, sume)

			if !used[to] {
				emap[to] = sume
				queue = append(queue, to)
				used[to] = true
			} else {
				if emap[to] < sume {
					emap[to] = sume
					queue = append(queue, to)
				}
			}
		}
	}

	return mx
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
