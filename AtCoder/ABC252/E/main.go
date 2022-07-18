// E - Road Reduction
// https://atcoder.jp/contests/abc252/tasks/abc252_e
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

type edge struct {
	idx  int
	to   int
	cost int
}

type intHeap [][]int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h intHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	abcs := make([]map[int]edge, n)
	for i := range abcs {
		abcs[i] = make(map[int]edge)
	}
	for i := 0; i < m; i++ {
		a, b, c := readInt()-1, readInt()-1, readInt()
		abcs[a][b] = edge{idx: i + 1, to: b, cost: c}
		abcs[b][a] = edge{idx: i + 1, to: a, cost: c}
	}

	queue := new(intHeap)
	heap.Init(queue)
	queue.Push([]int{0, 0})

	mincosts := make([]int, n)
	for i := range mincosts {
		mincosts[i] = -1
	}
	mincosts[0] = 0

	done := make([]bool, n)
	done[0] = true

	ans := make([]int, n)
	// ダイクストラ法
	for queue.Len() > 0 {
		ci := heap.Pop(queue).([]int)
		cost, idx := ci[0], ci[1]
		if cost != mincosts[idx] {
			continue
		}

		for _, eg := range abcs[idx] {
			sumCost := cost + eg.cost
			if mincosts[eg.to] < 0 || mincosts[eg.to] > sumCost {
				mincosts[eg.to] = sumCost
				ans[eg.to] = eg.idx
				queue.Push([]int{sumCost, eg.to})
			}
		}
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[3 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
