// RGB Coloring 2
// https://atcoder.jp/contests/abc199/tasks/abc199_d

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var used []bool
var list []int

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
		ab := abmap[a]
		abmap[a] = append(ab, b)
		ab = abmap[b]
		abmap[b] = append(ab, a)
	}

	used = make([]bool, n+1)
	cls := make([]int, n+1)
	ans := 1
	for i := 1; i <= n; i++ {
		if used[i] {
			continue
		}
		ans *= 3

		list = make([]int, 0)
		dfs(i, abmap)

		cls[list[0]] = 1
		ret := dfs2(1, abmap, cls)
		ans *= ret
	}
	return ans
}

func dfs(p int, abmap map[int][]int) {
	if used[p] {
		return
	}
	used[p] = true

	list = append(list, p)
	for _, ab := range abmap[p] {
		dfs(ab, abmap)
	}
}

func dfs2(i int, abmap map[int][]int, cls []int) int {
	if i >= len(list) {
		return 1
	}

	p := list[i]
	ret := 0
	for cl := 1; cl <= 3; cl++ {
		cls[p] = cl

		flg := false
		for _, d := range abmap[p] {
			if cls[d] == cl {
				flg = true
				break
			}
		}
		if flg {
			continue
		}

		ret += dfs2(i+1, abmap, cls)
	}

	cls[p] = 0
	return ret
}

func main() {
	fmt.Println(run(os.Stdin))
}
