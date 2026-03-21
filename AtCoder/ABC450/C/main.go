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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	ans := 0
	visited := make(map[[2]int]bool)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			result := dfs(i, j, ss, visited)
			if result == 1 {
				ans++
			}
		}
	}

	fmt.Fprint(out, ans)
}

func dfs(i, j int, ss []string, visited map[[2]int]bool) int {
	if visited[[2]int{i, j}] {
		return 0
	}
	if ss[i][j] == '#' {
		return 0
	}
	if i <= 0 || j <= 0 || i >= len(ss)-1 || j >= len(ss[0])-1 {
		return -1
	}

	visited[[2]int{i, j}] = true

	result := 1
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range dir {
		ni, nj := i+d[0], j+d[1]

		if dfs(ni, nj, ss, visited) == -1 {
			result = -1
		}
	}

	return result
}

func main() {
	run(os.Stdin, os.Stdout)
}
