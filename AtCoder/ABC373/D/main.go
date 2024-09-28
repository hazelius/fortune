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
	uvmap := make(map[int]map[int]int)
	for i := 0; i < m; i++ {
		u, v, w := readInt()-1, readInt()-1, readInt()
		wmap, ok := uvmap[u]
		if !ok {
			wmap = make(map[int]int)
		}
		wmap[v] = w
		uvmap[u] = wmap

		vmap, ok := uvmap[v]
		if !ok {
			vmap = make(map[int]int)
		}
		vmap[u] = -w
		uvmap[v] = vmap
	}

	ws := make([]int, n)
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		d(i, uvmap, used, ws)
	}

	str := fmt.Sprintf("%v", ws)
	fmt.Fprint(out, str[1:len(str)-1])
}

func d(a int, uvmap map[int]map[int]int, used []bool, ws []int) {
	que := []int{a}
	for len(que) > 0 {
		q := que[0]
		que = que[1:]
		wmap, ok := uvmap[q]
		if !ok {
			continue
		}
		used[q] = true
		for v, w := range wmap {
			if used[v] {
				continue
			}
			ws[v] = ws[q] + w
			que = append(que, v)
			used[v] = true
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
