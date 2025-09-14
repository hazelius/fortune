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

type event struct {
	t   int
	typ int
	c   cos
}

type cos struct {
	idx int
	t   int
	dur int
	cnt int
}

type events []event

func (h events) Len() int      { return len(h) }
func (h events) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h events) Less(i, j int) bool {
	if h[i].t == h[j].t {
		return h[i].c.idx < h[j].c.idx
	}
	return h[i].t < h[j].t
}
func (h *events) Push(x interface{}) { *h = append(*h, x.(event)) }
func (h *events) Pop() interface{} {
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

	n, k := readInt(), readInt()
	abcs := make([]cos, n)
	for i := range abcs {
		abcs[i] = cos{i, readInt(), readInt(), readInt()}
	}

	pque := new(events)
	heap.Init(pque)
	for _, abc := range abcs {
		heap.Push(pque, event{t: abc.t, typ: 1, c: abc})
	}
	machi := make([]cos, 0, n)

	seats := k
	for pque.Len() > 0 {
		e := heap.Pop(pque).(event)
		switch e.typ {
		case 1:
			if len(machi) == 0 && e.c.cnt <= seats {
				seats -= e.c.cnt
				fmt.Fprintln(out, e.t)
				heap.Push(pque, event{t: e.t + e.c.dur, typ: 2, c: e.c})
			} else {
				machi = append(machi, e.c)
			}
		case 2:
			seats += e.c.cnt
			for len(machi) > 0 {
				next := machi[0]
				if next.cnt <= seats {
					machi = machi[1:]
					seats -= next.cnt
					fmt.Fprintln(out, e.t)
					heap.Push(pque, event{t: e.t + next.dur, typ: 2, c: next})
				} else {
					break
				}
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
