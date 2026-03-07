package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var ans []bool
var as []int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as = make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	uvmap := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		u, v := readInt()-1, readInt()-1
		uvmap[u] = append(uvmap[u], v)
		uvmap[v] = append(uvmap[v], u)
	}

	ans = make([]bool, n)
	visited := make([]bool, n)
	history := make(map[int]bool)
	dfs(0, uvmap, visited, history)

	for _, v := range ans {
		if v {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func dfs(i int, uvmap map[int][]int, visited []bool, history map[int]bool) {
	if visited[i] {
		return
	}
	visited[i] = true
	flg := false
	if history[as[i]] == false {
		history[as[i]] = true
		flg = true
	} else {
		ans[i] = true
	}
	uvs := uvmap[i]
	for _, v := range uvs {
		if visited[v] {
			continue
		}
		if ans[i] {
			ans[v] = true
		}
		dfs(v, uvmap, visited, history)
	}
	visited[i] = false
	if flg {
		history[as[i]] = false
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
