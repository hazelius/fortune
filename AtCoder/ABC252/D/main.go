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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ans := n * (n - 1) * (n - 2) / 6

	mapCnt := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		cnt, ok := mapCnt[a]
		if !ok {
			cnt = 0
		}
		cnt++
		mapCnt[a] = cnt
	}

	for _, cnt := range mapCnt {
		if cnt > 2 {
			ans -= cnt * (cnt - 1) * (cnt - 2) / 6
		}
		if cnt > 1 {
			ans -= (cnt * (cnt - 1) / 2) * (n - cnt)
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
