package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type position struct {
	x int
	y int
}

type status struct {
	turn  int
	money int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ps := make([][]int, n)
	for i := range ps {
		ps[i] = make([]int, n)
		for j := range ps[i] {
			ps[i][j] = readInt()
		}
	}

	rs := make([][]int, n)
	for i := range rs {
		rs[i] = make([]int, n-1)
		for j := range rs[i] {
			rs[i][j] = readInt()
		}
	}

	ds := make([][]int, n-1)
	for i := range ds {
		ds[i] = make([]int, n)
		for j := range ds[i] {
			ds[i][j] = readInt()
		}
	}

	dp := make(map[position]map[int]status)
	ini := make(map[int]status)
	ini[0] = status{0, 0}
	dp[position{0, 0}] = ini

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			pos := position{x, y}

			stmap := dp[pos]
			for maxpay, st := range stmap {
				maxpay = max(maxpay, ps[pos.x][pos.y])

				nextArr := make([]position, 0)
				costArr := make([]int, 0)

				if pos.y < n-1 {
					nextArr = append(nextArr, position{pos.x, pos.y + 1})
					costArr = append(costArr, rs[pos.x][pos.y])
				}
				if pos.x < n-1 {
					nextArr = append(nextArr, position{pos.x + 1, pos.y})
					costArr = append(costArr, ds[pos.x][pos.y])
				}

				for idx, nextPos := range nextArr {
					nextCost := costArr[idx]
					nextSt := status{money: st.money, turn: st.turn}
					if nextSt.money < nextCost {
						cnt := (nextCost - nextSt.money) / maxpay
						if (nextCost-nextSt.money)%maxpay > 0 {
							cnt++
						}
						nextSt.money += maxpay * cnt
						nextSt.turn += cnt
					}
					nextSt.turn++
					nextSt.money -= nextCost
					oldStMap, ok := dp[nextPos]
					if !ok {
						dp[nextPos] = make(map[int]status)
					} else {
						oldSt, ok := oldStMap[maxpay]
						if ok {
							if oldSt.turn < nextSt.turn {
								continue
							}
							if oldSt.turn == nextSt.turn && oldSt.money >= nextSt.money {
								continue
							}
						}
					}
					dp[nextPos][maxpay] = nextSt
				}
			}
		}
	}

	ans := -1
	for _, v := range dp[position{n - 1, n - 1}] {
		if ans == -1 || ans > v.turn {
			ans = v.turn
		}
	}

	fmt.Fprint(out, ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
