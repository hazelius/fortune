package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt64() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

func run(n int64) int64 {
	var ans, i int64
	for i = 1; i < n; i++ {
		ans += i
		if ans >= n {
			return i
		}
	}
	return 1
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt64()
	fmt.Println(run(n))
}
