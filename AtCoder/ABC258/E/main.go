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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, q, x := readInt(), readInt(), readInt()
	ws := make([]int, n)
	for i := range ws {
		ws[i] = readInt()
	}

	wd := make([][]int, n)
	sum := 0
	idx := 0
	cnt := 0
	for i := range wd {
		for sum < x {
			sum += ws[idx]
			cnt++
			idx++
			if idx >= n {
				idx = 0
			}
		}
		wd[i] = []int{idx, cnt}
		sum -= ws[i]
		cnt--
	}

	boxmap := make(map[int]int)
	used := make(map[int]int)
	next := 0
	start := 0
	for i := 0; i < n; i++ {
		used[next] = i
		d, w := wd[next][0], wd[next][1]
		boxmap[i] = w
		if v, ok := used[d]; ok {
			start = v
			break
		}
		next = d

	}

	l := len(boxmap)
	ans := make([]int, q)
	for i := range ans {
		q := readInt() - 1
		if q < l {
			ans[i] = boxmap[q]
		} else {
			idx := q - start
			idx %= (l - start)
			ans[i] = boxmap[idx+start]
		}
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
