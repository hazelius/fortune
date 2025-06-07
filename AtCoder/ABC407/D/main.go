package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var h, w int
var ans int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	h, w = readInt(), readInt()
	hws := make([][]int, h)
	for i := range hws {
		hws[i] = make([]int, w)
		for j := range hws[i] {
			hws[i][j] = readInt()
		}
	}

	ans = 0
	used := make([]int, h*w)
	dfs(0, used, hws)

	fmt.Fprint(out, ans)
}

func dfs(i int, used []int, hws [][]int) {
	if i == h*w {
		tmp := 0
		for p, v := range used {
			if v == 1 {
				continue
			}
			tmp ^= hws[p/w][p%w]
		}
		if ans < tmp {
			ans = tmp
		}
		return
	}

	if used[i] != 0 {
		dfs(i+1, used, hws)
		return
	}
	used[i] = -1
	dfs(i+1, used, hws)

	used[i] = 1
	if (i+1)%w != 0 && used[i+1] == 0 {
		used[i+1] = 1
		dfs(i+1, used, hws)
		used[i+1] = 0
	}

	if i+w < len(used) {
		used[i+w] = 1
		dfs(i+1, used, hws)
		used[i+w] = 0
	}
	used[i] = 0
}

func main() {
	run(os.Stdin, os.Stdout)
}
