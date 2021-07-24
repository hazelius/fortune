package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	sc  *bufio.Scanner
	mod int = 1e9 + 7
)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	abmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		abmap[a] = append(abmap[a], b)
		abmap[b] = append(abmap[b], a)
	}

	used := make(map[int]bool)
	used[1] = true
	queue := make(map[int]int)
	for _, v := range abmap[1] {
		queue[v]++
	}
	return bfs(n, queue, abmap, used)
}

func bfs(dis int, queue map[int]int, abmap map[int][]int, used map[int]bool) int {
	ret := 0
	nextq := make(map[int]int)
	for a, cnt := range queue {
		if a == dis {
			ret = (ret + cnt) % mod
		}

		if ret > 0 {
			continue
		}

		for _, v := range abmap[a] {
			if used[v] {
				continue
			}
			nextq[v] = (nextq[v] + cnt) % mod
		}
	}

	if ret > 0 || len(nextq) == 0 {
		return ret
	}

	for a := range nextq {
		used[a] = true
	}
	return bfs(dis, nextq, abmap, used)
}

func main() {
	fmt.Println(run(os.Stdin))
}
