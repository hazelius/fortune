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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	abs := make([][]int, n)
	bcnt := make([]int, n)
	for i := 0; i < m; i++ {
		a, b := readInt()-1, readInt()-1
		abs[a] = append(abs[a], b)
		bcnt[b]++
	}

	pque := new(intHeap)
	heap.Init(pque)
	for b, cnt := range bcnt {
		if cnt == 0 {
			pque.Push(b)
		}
	}

	ans := []int{}
	for pque.Len() > 0 {
		a := heap.Pop(pque).(int)
		ans = append(ans, a+1)
		for _, b := range abs[a] {
			bcnt[b]--
			if bcnt[b] == 0 {
				heap.Push(pque, b)
			}
		}
	}

	if len(ans) < n {
		return "-1"
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
