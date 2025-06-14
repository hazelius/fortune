package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type node struct {
	p int
	w int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	abmap := make(map[int][]node)
	for i := 0; i < m; i++ {
		a, b, w := readInt()-1, readInt()-1, readInt()
		ns, ok := abmap[a]
		if !ok {
			ns = make([]node, 0)
		}
		ns = append(ns, node{b, w})
		abmap[a] = ns
	}

	ans := -1
	used := make(map[node]bool)
	used[node{0, 0}] = true
	que := abmap[0]
	for len(que) > 0 {
		q := que[0]
		que = que[1:]
		if used[q] {
			continue
		}
		used[q] = true
		if q.p == n-1 && (ans < 0 || ans > q.w) {
			ans = q.w
		}

		bs, ok := abmap[q.p]
		if !ok {
			continue
		}
		for _, b := range bs {
			que = append(que, node{b.p, b.w ^ q.w})
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
