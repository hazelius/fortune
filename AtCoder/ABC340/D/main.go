package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var dp []int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h intHeap) Less(i, j int) bool  { return dp[h[i]] < dp[h[j]] }
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

	n := readInt()
	abxs := make([][]int, n-1)
	for i := range abxs {
		abxs[i] = []int{readInt(), readInt(), readInt() - 1}
	}

	dp = make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
	}

	pque := new(intHeap)
	heap.Init(pque)
	pque.Push(0)

	for pque.Len() > 0 {
		stage := heap.Pop(pque).(int)
		abx := abxs[stage]
		a, b, x := abx[0], abx[1], abx[2]
		if dp[stage+1] > a+dp[stage] {
			dp[stage+1] = a + dp[stage]
			if stage+1 < n-1 {
				pque.Push(stage + 1)
			}
		}
		if dp[x] > b+dp[stage] {
			dp[x] = b + dp[stage]
			if x < n-1 {
				pque.Push(x)
			}
		}
	}

	fmt.Fprint(out, dp[n-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
