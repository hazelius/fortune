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

	q := readInt()
	ans := make([]int, 0)
	xmap := make(map[int]int)
	minque, maxque := new(intHeap), new(intHeap)
	heap.Init(minque)
	heap.Init(maxque)

	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x := readInt()
			_, ok := xmap[x]
			if !ok {
				heap.Push(minque, x)
				heap.Push(maxque, -x)
			}
			xmap[x]++

		case 2:
			x, c := readInt(), readInt()
			cnt, ok := xmap[x]
			if ok {
				cnt -= c
				if cnt < 0 {
					cnt = 0
				}
				xmap[x] = cnt
			}
		case 3:
			minx := heap.Pop(minque).(int)
			for {
				if cnt, ok := xmap[minx]; ok && cnt > 0 {
					heap.Push(minque, minx)
					break
				}
				minx = heap.Pop(minque).(int)
			}
			maxx := heap.Pop(maxque).(int) * -1
			for {
				if cnt, ok := xmap[maxx]; ok && cnt > 0 {
					heap.Push(maxque, -maxx)
					break
				}
				maxx = heap.Pop(maxque).(int) * -1
			}
			ans = append(ans, maxx-minx)
		}
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
