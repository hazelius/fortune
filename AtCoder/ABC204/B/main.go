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

func run(n int, as []int) int {
	ans := 0
	for _, a := range as {
		if a > 10 {
			ans += a - 10
		}
	}
	return ans
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
