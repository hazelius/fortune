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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	h, w, n := readInt(), readInt(), readInt()
	hmap := make(map[int][][]int)
	wmap := make(map[int][][]int)
	for i := 0; i < n; i++ {
		hw := []int{i, readInt(), readInt()}
		hmap[hw[1]] = append(hmap[hw[1]], hw)
		wmap[hw[2]] = append(wmap[hw[2]], hw)
	}

	ans := make([][]int, n)
	used := make([]bool, n)
	dfs(h, w, hmap, wmap, used, ans)
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

func dfs(h, w int, hmap, wmap map[int][][]int, used []bool, ans [][]int) {
	hws, ok := hmap[h]
	if ok {
		for _, hw := range hws {
			idx := hw[0]
			if used[idx] {
				continue
			}
			ans[idx] = []int{1, w - hw[2] + 1}
			used[idx] = true
			w -= hw[2]
		}
		delete(hmap, h)
		dfs(h, w, hmap, wmap, used, ans)
		return
	}

	hws, ok = wmap[w]
	if ok {
		for _, hw := range hws {
			idx := hw[0]
			if used[idx] {
				continue
			}
			ans[idx] = []int{h - hw[1] + 1, 1}
			used[idx] = true
			h -= hw[1]
		}
		delete(wmap, w)
		dfs(h, w, hmap, wmap, used, ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
