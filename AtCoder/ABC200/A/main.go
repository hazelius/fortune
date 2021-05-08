// https://atcoder.jp/contests/abc200/tasks/abc200_a
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

func run(n int) int {
	return ((n - 1) / 100) + 1
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	fmt.Println(run(n))
}
