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

	_, m := readInt(), readInt()
	as := make([]int, m)
	for i := range as {
		as[i] = readInt()
	}

	dp := make(map[int][]int)
	current := 1
	for i, a := range as {
		vals, ok := dp[a]
		vals2, ok2 := dp[a+1]
		if ok {
			dp[a+1] = vals
			if !ok2 {
				delete(dp, a)
			}
		}
		if ok2 {
			dp[a] = vals2
			if !ok {
				delete(dp, a+1)
			}
		}

		if current == a {
			current++
			dp[a] = append(dp[a], i+1)
		} else if current == a+1 {
			current = a
			dp[a+1] = append(dp[a+1], i+1)
		}
	}

	mapa := make(map[int]int)
	for k, v := range dp {
		for _, a := range v {
			mapa[a] = k
		}
	}

	for i := 1; i <= m; i++ {
		ans, ok := mapa[i]
		if ok {
			fmt.Fprintln(out, ans)
		} else {
			fmt.Fprintln(out, current)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
