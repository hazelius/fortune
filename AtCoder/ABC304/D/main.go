package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type piece struct {
	x int
	y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	w, h := readInt(), readInt()
	n := readInt()
	pqs := make([][]int, n)
	for i := range pqs {
		pqs[i] = []int{readInt(), readInt()}
	}

	a := readInt()
	as := make([]int, a+2)
	as[0] = 0
	for i := 1; i <= a; i++ {
		as[i] = readInt()
	}
	as[a+1] = w

	b := readInt()
	bs := make([]int, b+2)
	bs[0] = 0
	for i := 1; i <= b; i++ {
		bs[i] = readInt()
	}
	bs[b+1] = h

	ans := make(map[piece]int)

	for _, pq := range pqs {
		x := sort.Search(a+1, func(i int) bool { return as[i] > pq[0] })
		y := sort.Search(b+1, func(i int) bool { return bs[i] > pq[1] })
		pie := piece{x, y}
		v := ans[pie]
		v++
		ans[pie] = v
	}

	min := n
	max := 0
	for _, cnt := range ans {
		if cnt > max {
			max = cnt
		}
		if cnt < min {
			min = cnt
		}
	}

	if len(ans) < (a+1)*(b+1) {
		min = 0
	}

	fmt.Fprintf(out, "%v %v", min, max)
}

func main() {
	run(os.Stdin, os.Stdout)
}
