// https://atcoder.jp/contests/abc208/tasks/abc208_d
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var inf int = 1 << 60

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	roads := make([][]int, n+1)
	for i := range roads {
		roads[i] = make([]int, n+1)
		for j := range roads[i] {
			if i == j {
				continue
			}
			roads[i][j] = inf
		}
	}

	for i := 0; i < m; i++ {
		roads[readInt()][readInt()] = readInt()
	}

	ans := 0
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			if roads[i] == nil {
				roads[i] = make([]int, n+1)
			}
			for j := 1; j <= n; j++ {
				roads[i][j] = min(roads[i][j], roads[i][k]+roads[k][j])
				if roads[i][j] >= inf {
					continue
				}
				ans += roads[i][j]
			}
		}
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
