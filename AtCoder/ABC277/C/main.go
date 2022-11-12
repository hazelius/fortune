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
var used map[int]bool

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	amaps := make(map[int][]int)
	for i := 0; i < n; i++ {
		a, b := readInt(), readInt()
		amaps[a] = append(amaps[a], b)
		amaps[b] = append(amaps[b], a)
	}

	ans = 0
	used = make(map[int]bool)
	dfs(1, amaps)

	fmt.Fprint(out, ans)
}

func dfs(now int, amaps map[int][]int) {
	if used[now] {
		return
	}
	used[now] = true

	if now > ans {
		ans = now
	}

	ts := amaps[now]
	for _, t := range ts {
		dfs(t, amaps)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
