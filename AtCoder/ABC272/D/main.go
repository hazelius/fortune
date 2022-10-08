package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var ans [][]int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()

	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}

	dirs := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d := i*i + j*j
			if m == d {
				dirs = append(dirs, []int{i, j})
				if i != 0 {
					dirs = append(dirs, []int{-i, j})
				}
				if j != 0 {
					dirs = append(dirs, []int{i, -j})
				}
				if i != 0 && j != 0 {
					dirs = append(dirs, []int{-i, -j})
				}
			}
		}
	}

	q := [][]int{{0, 0}}
	ans[0][0] = 0
	for len(q) > 0 {
		q = f(q, dirs)
	}

	for _, row := range ans {
		strRow := fmt.Sprintf("%v", row)
		fmt.Fprintln(out, strRow[1:len(strRow)-1])
	}
}

func f(q, dirs [][]int) [][]int {
	if len(q) == 0 {
		return q
	}
	n := len(ans)
	xy := q[0]
	q = q[1:]
	cnt := ans[xy[0]][xy[1]] + 1

	for _, dir := range dirs {
		x := dir[0] + xy[0]
		y := dir[1] + xy[1]

		if x < 0 || y < 0 || x >= n || y >= n {
			continue
		}
		if ans[x][y] < 0 {
			ans[x][y] = cnt
			q = append(q, []int{x, y})
		}
	}

	return q
}

func main() {
	run(os.Stdin, os.Stdout)
}
