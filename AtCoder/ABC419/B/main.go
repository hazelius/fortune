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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	q := readInt()

	queue := new(intHeap)
	heap.Init(queue)

	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x := readInt()
			heap.Push(queue, x)
		case 2:
			ans := heap.Pop(queue).(int)
			fmt.Fprintln(out, ans)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
