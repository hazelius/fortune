// E - Stronger Takahashi
// https://atcoder.jp/contests/abc213/tasks/abc213_e
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

type queue struct {
	rgt    int
	lft    int
	values [][]int
}

func newQueue(l int) queue {
	q := queue{rgt: l, lft: l}
	q.values = make([][]int, l*8)
	return q
}

func (q *queue) pop() []int {
	q.rgt++
	return q.values[q.rgt-1]
}

func (q *queue) pushfront(val []int) {
	q.rgt--
	q.values[q.rgt] = val
}

func (q *queue) pushback(val []int) {
	q.values[q.lft] = val
	q.lft++
}

func (q *queue) length() int {
	return q.lft - q.rgt
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, w)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][0] = 0
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	deque := newQueue(h * w)
	deque.pushback([]int{0, 0})

	for deque.length() > 0 {
		xy := deque.pop()
		x, y := xy[0], xy[1]
		if visited[x][y] {
			continue
		}
		visited[x][y] = true
		cost := dp[x][y]
		next := [][]int{{x - 1, y}, {x, y - 1}, {x + 1, y}, {x, y + 1}}
		for _, np := range next {
			nx, ny := np[0], np[1]
			if nx < 0 || ny < 0 || nx >= h || ny >= w {
				continue
			}
			if ss[nx][ny] == '#' {
				continue
			}
			if dp[nx][ny] <= cost {
				continue
			}
			dp[nx][ny] = cost
			deque.pushfront([]int{nx, ny})
		}

		for i := -2; i <= 2; i++ {
			for j := -2; j <= 2; j++ {
				if abs(i)+abs(j) > 3 {
					continue
				}
				nx, ny := x+i, y+j
				if nx < 0 || ny < 0 || nx >= h || ny >= w {
					continue
				}
				if dp[nx][ny] <= cost+1 {
					continue
				}
				dp[nx][ny] = cost + 1
				deque.pushback([]int{nx, ny})
			}
		}
	}

	return dp[h-1][w-1]
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
