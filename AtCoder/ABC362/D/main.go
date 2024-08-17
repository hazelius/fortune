package main

import (
	"bufio"
	"container/heap"
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

type node struct {
	idx  int
	cost int
}

type nodes []node

func (h nodes) Len() int            { return len(h) }
func (h nodes) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h nodes) Less(i, j int) bool  { return h[i].cost < h[j].cost }
func (h *nodes) Push(x interface{}) { *h = append(*h, x.(node)) }
func (h *nodes) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	uvbmap := make(map[int]map[int]int)
	for i := 0; i < m; i++ {
		u, v, b := readInt()-1, readInt()-1, readInt()
		if _, ok := uvbmap[u]; !ok {
			uvbmap[u] = make(map[int]int)
		}
		if _, ok := uvbmap[v]; !ok {
			uvbmap[v] = make(map[int]int)
		}
		uvbmap[u][v] = b
		uvbmap[v][u] = b
	}

	ws := make([]int, n)
	for i := range ws {
		ws[i] = -1
	}
	ws[0] = as[0]

	pque := new(nodes)
	heap.Init(pque)
	pque.Push(node{0, as[0]})

	// ダイクストラ法
	for pque.Len() > 0 {
		cur := heap.Pop(pque).(node)
		if cur.cost != ws[cur.idx] {
			continue
		}
		uvb, ok := uvbmap[cur.idx]
		if !ok {
			break
		}

		for v, b := range uvb {
			vw := cur.cost + b + as[v]
			if ws[v] < 0 || ws[v] > vw {
				ws[v] = vw
				pque.Push(node{v, vw})
			}
		}
	}

	ans := fmt.Sprintf("%v", ws[1:])
	fmt.Fprint(out, ans[1:len(ans)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
