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

type mas struct {
	x  int
	y  int
	st int
}

type intHeap []mas

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h intHeap) Less(i, j int) bool  { return h[i].st < h[j].st }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(mas)) }
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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	h, w, x := readInt(), readInt(), readInt()
	p, q := readInt()-1, readInt()-1
	ss := make([][]int, h)
	for i := range ss {
		ss[i] = make([]int, w)
		for j := range ss[i] {
			ss[i][j] = readInt()
		}
	}

	queue := new(intHeap)
	heap.Init(queue)
	queue.Push(mas{p, q, ss[p][q]})

	t := 0
	used := make([][]bool, h)
	for i := range used {
		used[i] = make([]bool, w)
	}
	used[p][q] = true

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for queue.Len() > 0 {
		ci := heap.Pop(queue).(mas)
		// ci.st*x >= t とすると64bit整数の範囲を超えて正しく判定できないので割り算で判定する
		val := t / x
		if t%x == 0 {
			val--
		}

		if !(ci.x == p && ci.y == q) && ci.st > val {
			break
		}

		t += ci.st

		for _, dir := range dirs {
			ni, nj := ci.x+dir[0], ci.y+dir[1]
			if ni < 0 || nj < 0 || ni >= h || nj >= w {
				continue
			}
			if used[ni][nj] {
				continue
			}
			used[ni][nj] = true

			heap.Push(queue, mas{ni, nj, ss[ni][nj]})
		}
	}

	fmt.Fprint(out, t)
}

func main() {
	run(os.Stdin, os.Stdout)
}
