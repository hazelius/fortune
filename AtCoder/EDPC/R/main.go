// R - Walk
// https://atcoder.jp/contests/dp/tasks/dp_r
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const mod = 1e9 + 7

type matrix struct {
}

func (m *matrix) pow(k int) {

}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n, k int, as [][]bool) int {
	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	as := make([][]bool, n)
	for i := range as {
		as[i] = make([]bool, n)
		for j := range as[i] {
			as[i][j] = (readInt() == 1)
		}
	}
	fmt.Println(run(n, k, as))
}
