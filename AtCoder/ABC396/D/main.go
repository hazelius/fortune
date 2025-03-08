package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var ans int
var n int
var m int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m = readInt(), readInt()
	uvmap := make(map[int]map[int]int)
	for i := 0; i < m; i++ {
		u, v, w := readInt()-1, readInt()-1, readInt()
		if _, ok := uvmap[u]; !ok {
			uvmap[u] = make(map[int]int)
		}
		if _, ok := uvmap[v]; !ok {
			uvmap[v] = make(map[int]int)
		}
		uvmap[u][v] = w
		uvmap[v][u] = w
	}

	ans = 1 << 60
	used := make([]bool, n)
	d(0, uvmap, used, 0)

	fmt.Fprint(out, ans)
}

func d(a int, uvmap map[int]map[int]int, used []bool, w int) {
	if a == n-1 {
		if ans > w {
			ans = w
		}
		return
	}

	used[a] = true
	usmap, ok := uvmap[a]
	if !ok {
		return
	}
	for v, label := range usmap {
		if used[v] {
			continue
		}
		d(v, uvmap, used, w^label)
	}

	used[a] = false
}

func main() {
	run(os.Stdin, os.Stdout)
}
