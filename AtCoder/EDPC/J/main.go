// J - Sushi
// https://atcoder.jp/contests/dp/tasks/dp_j
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var dp map[int]float64

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, as []int) float64 {
	var a1, a2, a3 int
	for _, v := range as {
		switch v {
		case 1:
			a1++
		case 2:
			a2++
		case 3:
			a3++
		}
	}

	dp = make(map[int]float64)
	dp[1000000] = float64(n)
	return f(a1, a2, a3, float64(n))
}

func f(a1, a2, a3 int, n float64) float64 {
	if a1 < 0 || a2 < 0 || a3 < 0 {
		return 0
	}

	key := a1*1000000 + a2*1000 + a3
	if v, ok := dp[key]; ok {
		return v
	}

	ret := n
	ret += f(a1-1, a2, a3, n) * float64(a1)
	ret += f(a1+1, a2-1, a3, n) * float64(a2)
	ret += f(a1, a2+1, a3-1, n) * float64(a3)
	ret /= float64(a1 + a2 + a3)

	dp[key] = ret
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, as))
}
