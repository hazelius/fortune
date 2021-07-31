package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

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

	q := readInt()
	q2 := 0
	que := new(intHeap)
	ans := make([]string, 0)
	heap.Init(que)

	for i := 0; i < q; i++ {
		query := readInt()
		switch query {
		case 1:
			x := readInt()
			heap.Push(que, x-q2)
		case 2:
			q2 += readInt()
		case 3:
			ret := heap.Pop(que).(int) + q2
			ans = append(ans, strconv.Itoa(ret))
		}
	}

	return strings.Join(ans, "\n")
}

func main() {
	fmt.Println(run(os.Stdin))
}
