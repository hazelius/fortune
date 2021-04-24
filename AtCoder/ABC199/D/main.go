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
	mapab := make(map[int][]int)
	for _, ab := range abs {
		a, b := ab[0], ab[1]
		mapab[a] = append(mapab[a], b)
		mapab[b] = append(mapab[b], a)
	}

	dp := make([]int, n+1)
	dp[1] = 3
	ans := 3
	for i := 2; i <= n; i++ {
		ans *= f(i, 0, dp, mapab)
	}
	fmt.Println(dp)
	return ans
}

func f(c, p int, dp []int, mapab map[int][]int) int {
	if dp[c] > 0 {
		return 1
	}

	s, ok := mapab[c]
	if !ok {
		return 3
	}

	for _, v := range s {
		if v == p {
			continue
		}
		if dp[v] > 0 {
			n--
		} else {
		}
		f(v, c, dp, mapab)
	}

	return dp[c][0] + dp[c][1] + dp[c][2]
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
