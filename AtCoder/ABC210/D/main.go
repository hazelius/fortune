package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var (
	as      [][]int
	c       int
	ans     int
	mincost int
)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	c = readInt()
	as = make([][]int, h)
	mincost = 1e9
	for i := range as {
		as[i] = make([]int, w)
		for j := range as[i] {
			as[i][j] = readInt()
			if as[i][j] < mincost {
				mincost = as[i][j]
			}
		}
	}

	ans = as[0][0] + as[0][1] + c
	for i, row := range as {
		for j := range row {
			dfs(i, j, i, j)
		}
	}

	return ans
}

func dfs(x1, y1, x2, y2 int) {
	dis := abs(x2-x1) + abs(y2-y1)
	if dis > 0 {
		if ans <= c*dis+as[x1][y1]+mincost {
			return
		}

		cost := as[x1][y1] + as[x2][y2] + dis*c
		if ans > cost {
			ans = cost
		}
	}

	if y1 <= y2 && y2+1 < len(as[0]) {
		dfs(x1, y1, x2, y2+1)
	}
	if x1 != x2 && y1 >= y2 && y2 > 0 {
		dfs(x1, y1, x2, y2-1)
	}
	if y1 == y2 && x2+1 < len(as) {
		dfs(x1, y1, x2+1, y2)
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return a * -1
}

func main() {
	fmt.Println(run(os.Stdin))
}
