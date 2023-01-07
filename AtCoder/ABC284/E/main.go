package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var ans int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	_, m := readInt(), readInt()
	uvmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		u, v := readInt(), readInt()
		us, ok := uvmap[u]
		if !ok {
			us = make([]int, 0)
		}
		us = append(us, v)
		uvmap[u] = us

		vs, ok := uvmap[v]
		if !ok {
			vs = make([]int, 0)
		}
		vs = append(vs, u)
		uvmap[v] = vs
	}

	used := make(map[int]bool)
	ans = 0
	dfs(1, uvmap, used)

	fmt.Fprint(out, ans)
}

func dfs(a int, uvmap map[int][]int, used map[int]bool) {
	if used[a] {
		return
	}
	used[a] = true

	if ans >= 1000000 {
		return
	}
	ans++

	us := uvmap[a]
	for _, u := range us {
		dfs(u, uvmap, used)
	}
	used[a] = false
}

func main() {
	run(os.Stdin, os.Stdout)
}
