package main

import (
	"bufio"
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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	ps := make([]int, n-1)
	for i := range ps {
		ps[i] = readInt() - 1
	}

	xys := make([]int, n)
	for i := 0; i < m; i++ {
		x, y := readInt()-1, readInt()+1
		if xys[x] < y {
			xys[x] = y
		}
	}

	parents := make(map[int]map[int]bool)
	for i, p := range ps {
		if _, ok := parents[p]; !ok {
			parents[p] = make(map[int]bool)
		}
		parents[p][i+1] = true
	}

	cnt := 0
	que := [][]int{{0, xys[0]}}
	for len(que) > 0 {
		p, ins := que[0][0], que[0][1]
		que = que[1:]
		if ins < xys[p] {
			ins = xys[p]
		}
		if ins > 0 {
			cnt++
		}

		for np := range parents[p] {
			que = append(que, []int{np, ins - 1})
		}
	}

	fmt.Fprint(out, cnt)
}

func main() {
	run(os.Stdin, os.Stdout)
}
