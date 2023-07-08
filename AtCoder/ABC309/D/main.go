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

	n1, n2, m := readInt(), readInt(), readInt()

	abmap := make(map[int]map[int]bool)
	for i := 0; i < m; i++ {
		a, b := readInt()-1, readInt()-1
		if _, ok := abmap[a]; !ok {
			abmap[a] = map[int]bool{}
		}
		if _, ok := abmap[b]; !ok {
			abmap[b] = map[int]bool{}
		}
		abmap[a][b] = true
		abmap[b][a] = true
	}

	ans1 := fds(0, abmap)
	ans2 := fds(n1+n2-1, abmap)

	fmt.Fprint(out, ans1+ans2+1)
}

func fds(i int, abmap map[int]map[int]bool) int {
	que := [][]int{{i, 0}}
	used := make(map[int]bool)
	ret := 0
	for len(que) > 0 {
		q, cnt := que[0][0], que[0][1]
		que = que[1:]
		if used[q] {
			continue
		}
		used[q] = true
		if ret < cnt {
			ret = cnt
		}

		for p := range abmap[q] {
			if used[p] {
				continue
			}
			que = append(que, []int{p, cnt + 1})
		}
	}
	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
