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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	as := make([][]int, h)
	for i := range as {
		as[i] = make([]int, w)
		for j := range as[i] {
			as[i][j] = readInt()
		}
	}

	ans = 0
	used := make(map[int]bool)

	dfs(0, 0, used, as)

	fmt.Fprint(out, ans)
}

func dfs(x, y int, used map[int]bool, as [][]int) {
	a := as[x][y]
	if used[a] {
		return
	}
	if x == len(as)-1 && y == len(as[0])-1 {
		ans++
		return
	}

	used[a] = true
	if x < len(as)-1 {
		dfs(x+1, y, used, as)
	}
	if y < len(as[0])-1 {
		dfs(x, y+1, used, as)
	}
	used[a] = false
}

func main() {
	run(os.Stdin, os.Stdout)
}
