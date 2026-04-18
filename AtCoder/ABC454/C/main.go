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
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	abs := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		abs[i] = make([]int, 0)
	}
	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		abs[a] = append(abs[a], b)
	}
	ansmap := make(map[int]bool)
	ansmap[1] = true
	dfs(1, abs, ansmap)
	fmt.Fprint(out, len(ansmap))
}

func dfs(current int, abs [][]int, ansmap map[int]bool) {
	for _, next := range abs[current] {
		if !ansmap[next] {
			ansmap[next] = true
			dfs(next, abs, ansmap)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
