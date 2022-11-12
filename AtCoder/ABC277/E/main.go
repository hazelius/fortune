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
	p   int
	btn int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m, k := readInt(), readInt(), readInt()
	uvamaps := make(map[int][]node)
	for i := 0; i < m; i++ {
		u, v, a := readInt(), readInt(), readInt()
		uvamaps[u] = append(uvamaps[u], node{p: v, btn: a})
		uvamaps[v] = append(uvamaps[v], node{p: u, btn: a})
	}

	smap := make(map[int]bool)
	for i := 0; i < k; i++ {
		s := readInt()
		smap[s] = true
	}

	used := make(map[node]bool)

	que := [][]int{{1, 0, 0}}
	for {
		if len(que) == 0 {
			break
		}
		q := que[0]
		now := node{q[0], q[1]}
		cnt := q[2]
		que = que[1:]

		if used[now] {
			continue
		}
		used[now] = true

		if now.p == n {
			fmt.Fprint(out, cnt)
			return
		}

		vas, ok := uvamaps[now.p]
		if !ok {
			continue
		}

		cnt++
		for _, va := range vas {
			if va.btn == now.btn {
				continue
			}
			next := node{p: va.p, btn: now.btn}
			if used[next] {
				continue
			}
			que = append(que, []int{next.p, next.btn, cnt})
		}

		if smap[now.p] {
			for _, va := range vas {
				nb := 1 - now.btn
				if va.btn == nb {
					continue
				}
				next := node{p: va.p, btn: nb}
				if used[next] {
					continue
				}
				que = append(que, []int{next.p, next.btn, cnt})
			}
		}
	}

	fmt.Fprint(out, -1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
