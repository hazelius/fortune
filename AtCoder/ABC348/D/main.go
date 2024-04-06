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

type Point struct {
	x int
	y int
}

type Pointe struct {
	p Point
	e int
}

type Pointes []Pointe

func (h Pointes) Len() int            { return len(h) }
func (h Pointes) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h Pointes) Less(i, j int) bool  { return h[i].e > h[j].e }
func (h *Pointes) Push(x interface{}) { *h = append(*h, x.(Pointe)) }
func (h *Pointes) Pop() interface{} {
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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}
	n := readInt()
	rcemap := make(map[Point]int)
	for i := 0; i < n; i++ {
		x, y, e := readInt()-1, readInt()-1, readInt()
		rcemap[Point{x, y}] = e
	}

	ques := new(Pointes)
	heap.Init(ques)
	for i, s := range ss {
		idx := strings.Index(s, "S")
		if idx >= 0 {
			ques.Push(Pointe{Point{i, idx}, 0})
			break
		}
	}

	used := make(map[Point]int)
	dis := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for ques.Len() > 0 {
		pt := heap.Pop(ques).(Pointe)
		e, ok := used[pt.p]
		if ok && e >= pt.e {
			continue
		}
		used[pt.p] = pt.e

		if e, ok := rcemap[pt.p]; ok {
			if pt.e < e {
				pt.e = e
			}
		}

		if pt.e <= 0 {
			continue
		}

		pt.e--
		for _, d := range dis {
			x := pt.p.x + d[0]
			y := pt.p.y + d[1]
			if x < 0 || y < 0 || x >= h || y >= w {
				continue
			}
			c := ss[x][y]
			if c == '#' {
				continue
			} else if c == 'T' {
				fmt.Fprint(out, "Yes")
				return
			}
			newp := Point{x, y}
			e, ok := used[newp]
			if ok && e >= pt.e {
				continue
			}
			heap.Push(ques, Pointe{newp, pt.e})
		}
	}

	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
