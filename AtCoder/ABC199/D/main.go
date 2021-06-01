package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n, m int, abs [][]int) int {
	sides := make([][]int, n+1)
	for _, ab := range abs {
		a, b := ab[0], ab[1]
		sides[a] = append(sides[a], b)
		sides[b] = append(sides[b], a)
	}

	dp := make([][]int, n+1)
	ans := 0
	for i := 2; i <= n; i++ {
		dfs(i, 0, dp, sides)
	}
	return ans
}

func dfs(c, p int, dp, sides [][]int) {
	if dp[c] != nil {
		return
	}

	dp[c] = make([]int, 0)

	for _, v := range sides[c] {
		dfs(v, c, dp, sides)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := readInt(), readInt()
	abs := make([][]int, m)
	for i := range abs {
		abs[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(n, m, abs))
}
