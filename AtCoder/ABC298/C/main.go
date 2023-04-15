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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()

	boxes := make([]*intHeap, n)
	for i := range boxes {
		boxes[i] = new(intHeap)
		heap.Init(boxes[i])
	}

	nums := make(map[int]*intHeap)

	for cnt := 0; cnt < q; cnt++ {
		switch readInt() {
		case 1:
			i, j := readInt(), readInt()-1
			heap.Push(boxes[j], i)

			nn, ok := nums[i]
			if !ok {
				nn = new(intHeap)
			}
			heap.Push(nn, j+1)
			nums[i] = nn
		case 2:
			i := readInt() - 1
			tmp := new(intHeap)
			heap.Init(tmp)
			for boxes[i].Len() > 0 {
				p := heap.Pop(boxes[i]).(int)
				heap.Push(tmp, p)
				fmt.Fprint(out, p)
				fmt.Fprint(out, " ")
			}
			fmt.Fprintln(out, "")
			boxes[i] = tmp
		case 3:
			i := readInt()
			tmp := new(intHeap)
			heap.Init(tmp)
			pre := -1
			for nums[i].Len() > 0 {
				p := heap.Pop(nums[i]).(int)
				if p == pre {
					continue
				}
				pre = p
				heap.Push(tmp, p)
				fmt.Fprint(out, p)
				fmt.Fprint(out, " ")
			}
			fmt.Fprintln(out, "")
			nums[i] = tmp
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
