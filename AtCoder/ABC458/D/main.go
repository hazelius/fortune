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

type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}

type lheap []int

func (h lheap) Len() int            { return len(h) }
func (h lheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h lheap) Less(i, j int) bool  { return h[i] > h[j] }
func (h *lheap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *lheap) Pop() interface{} {
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

	x := readInt()
	q := readInt()

	lque := new(lheap)
	heap.Init(lque)
	rque := new(intHeap)
	heap.Init(rque)

	for i := 0; i < q; i++ {
		a, b := readInt(), readInt()
		if a > b {
			a, b = b, a
		}
		if a <= x && x <= b {
			heap.Push(lque, a)
			heap.Push(rque, b)
			fmt.Fprintln(out, x)
			continue
		}
		if x < a {
			heap.Push(rque, a)
			heap.Push(rque, b)
			heap.Push(lque, x)
			x = heap.Pop(rque).(int)
			fmt.Fprintln(out, x)
			continue
		}
		if b < x {
			heap.Push(lque, a)
			heap.Push(lque, b)
			heap.Push(rque, x)
			x = heap.Pop(lque).(int)
			fmt.Fprintln(out, x)
			continue
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
