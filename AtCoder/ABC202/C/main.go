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

func run(n int, as, bs, cs []int) int {
	bc := make(map[int]int)
	for _, c := range cs {
		val := bs[c-1]
		if _, ok := bc[val]; ok {
			bc[val]++
		} else {
			bc[val] = 1
		}
	}

	ans := 0
	for _, v := range as {
		ans += bc[v]
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	as, bs, cs := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = readInt()
	}
	for i := 0; i < n; i++ {
		bs[i] = readInt()
	}
	for i := 0; i < n; i++ {
		cs[i] = readInt()
	}
	fmt.Println(run(n, as, bs, cs))
}
