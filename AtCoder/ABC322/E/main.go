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

	n, k, p := readInt(), readInt(), readInt()
	costs := make([]int, n)
	as := make([][]int, n)
	for i := range as {
		costs[i] = readInt()
		as[i] = make([]int, k)
		for j := range as[i] {
			as[i][j] = readInt()
		}
	}

	dp := make([]map[int]int, n+1)
	dp[0] = make(map[int]int)
	dp[0][0] = 0

	for i, a := range as {
		dp[i+1] = make(map[int]int)
		for key, cost := range dp[i] {
			precost, ok := dp[i+1][key]
			if !ok {
				dp[i+1][key] = cost
			} else if precost > cost {
				dp[i+1][key] = cost
			}

			params := make([]int, k)
			for j := 0; j < k; j++ {
				params[j] = key % 10
				key /= 10
			}
			for j, par := range a {
				params[j] += par
				if params[j] > p {
					params[j] = p
				}
			}

			jo := 1
			w := 0
			for _, prm := range params {
				w += prm * jo
				jo *= 10
			}

			newCost := cost + costs[i]
			mincost, ok := dp[i+1][w]
			if !ok {
				dp[i+1][w] = newCost
			} else if mincost > newCost {
				dp[i+1][w] = newCost
			}
		}
	}

	w := p
	for i := 1; i < k; i++ {
		w *= 10
		w += p
	}

	ans, ok := dp[n][w]
	if !ok {
		ans = -1
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
