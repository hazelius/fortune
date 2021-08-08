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

var used map[int]bool
var ans []int

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	abmap := make(map[int]*intHeap, n)
	for i := 0; i < n; i++ {
		a, b := readInt(), readInt()
		if _, ok := abmap[a]; !ok {
			abmap[a] = new(intHeap)
			heap.Init(abmap[a])
		}
		if _, ok := abmap[b]; !ok {
			abmap[b] = new(intHeap)
			heap.Init(abmap[b])
		}
		heap.Push(abmap[a], b)
		heap.Push(abmap[b], a)
	}

	used = make(map[int]bool)
	ans = make([]int, 0, n)
	dfs(1, abmap)

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func dfs(n int, abmap map[int]*intHeap) bool {
	if used[n] {
		return false
	}
	used[n] = true
	ans = append(ans, n)

	as := abmap[n]
	for as.Len() > 0 {
		b := heap.Pop(as).(int)
		if used[b] {
			continue
		}
		if dfs(b, abmap) {
			ans = append(ans, n)
		}
	}
	return true
}

func main() {
	fmt.Println(run(os.Stdin))
}
